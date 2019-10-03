package components

import (
	"cin/base"
	"cin/configs"
	"cin/controllers"
	"cin/models"
	"cin/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type WebsocketServer struct {
	Base

	port uint16 // websocket 监听端口

	upgrader          websocket.Upgrader         // websocket http 升级为 websocket 的方法
	handlerDict       map[string]reflect.Type    // 控制器反射map
	handlerActionDict map[string]map[string]bool // 控制器 => 方法 => 是否存在（注册时记录）

	onConnectHandler func(conn *models.WebsocketSession)                     // 自定义方法：有新连接连入
	onMessageHandler func(conn *models.WebsocketSession, recvMessage []byte) // 自定义方法：收到消息
	onCloseHandler   func(conn *models.WebsocketSession)                     // 自定义方法：连接被关闭
}

// 使用配置 初始化数据
func (component *WebsocketServer) Init(configInterface base.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *configs.WebsocketServer
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*configs.WebsocketServer)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(configs.WebsocketServer)
		config = &configStruct
	} else {
		panic("illegal config")
	}


	component.port = config.Port
	component.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	component.handlerDict = map[string]reflect.Type{}
	component.handlerActionDict = map[string]map[string]bool{}
	component.onConnectHandler = nil
	component.onMessageHandler = nil
	component.onCloseHandler = nil

	// 注册处理器（控制器）
	for _, handler := range config.HandlerList {
		component.Register(handler)
	}
	component.onMessageHandler = config.OnWebsocketMessageHandler
}

// 开始
func (component *WebsocketServer) Start() {
	component.Base.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.port), 10),
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// 注册处理器。注册后将根据路由自动匹配调用方法
func (component *WebsocketServer) Register(handlerInterface controllers.HandlerInterface) {
	t := reflect.TypeOf(handlerInterface)
	handlerType := t.Elem() // 获取实体
	handlerName := handlerType.Name()
	handlerName = strings.TrimSuffix(handlerName, "Handler")
	handlerName = strings.TrimSuffix(handlerName, "Controller")
	handlerName = strings.ToLower(handlerName)
	component.handlerDict[handlerName] = t

	// 保存控制器啊所有方法名字
	component.handlerActionDict[handlerName] = map[string]bool{}
	methodLen := handlerType.NumMethod()
	for i := 0; i < methodLen; i++ {
		method := handlerType.Method(i)
		methodName := method.Name
		component.handlerActionDict[handlerName][methodName] = true
	}
}

// 设置 接受新连接的方法
func (component *WebsocketServer) OnConnect(handler func(conn *models.WebsocketSession)) {
	component.onConnectHandler = handler
}

// 设置 接受消息的方法
func (component *WebsocketServer) OnMessage(handler func(conn *models.WebsocketSession, recvMessage []byte)) {
	component.onMessageHandler = handler
}

// 设置 关闭连接的方法
func (component *WebsocketServer) OnClose(handler func(conn *models.WebsocketSession)) {
	component.onCloseHandler = handler
}

// 处理收到消息的方法
func (component *WebsocketServer) handlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := component.upgrader.Upgrade(w, r, nil)
	if conn == nil || err != nil {
		return
	}
	session := models.NewWebsocketSession(conn)
	component.callOnConnect(session)

	defer func() {
		component.callOnClose(session)
		_ = session.Close()
	}()

	for {
		var recvMessage []byte
		var messageType int
		messageType, recvMessage, err = conn.ReadMessage()
		if err != nil {
			break
		}
		// 常规消息处理
		if messageType == websocket.TextMessage || messageType == websocket.BinaryMessage {
			component.callOnMessage(session, recvMessage)
			// 处理客户端发送的请求并返回数据。如果返回 nil ，则代表不会返回数据给客户端
			sendMessage := component.getSendMessageByHandler(session, recvMessage)
			if sendMessage != nil {
				err = conn.WriteMessage(websocket.TextMessage, sendMessage)
				if err != nil {
				}
			}
		}
	}
}

// 使用处理器获取返回数据
func (component *WebsocketServer) getSendMessageByHandler(session *models.WebsocketSession, recvMessage []byte) (sendMessage []byte) {
	defer func() {
		if rec := recover(); rec != nil {
			sendMessage = utils.Json.Encode(rec)
		}
	}()
	session.Send(nil)

	// 使用自定义方法处理
	component.callOnMessage(session, recvMessage)
	sendMessage = session.GetSendMessage()
	if sendMessage != nil {
		return sendMessage
	}

	// 如果没有数据返回，则使用处理器尝试匹配路由
	sendMessage = component.callHandler(session, recvMessage)

	return sendMessage
}

// 获取返回结果
func (component *WebsocketServer) callHandler(session *models.WebsocketSession, recvMessage []byte) []byte {
	recvMessageModel := new(models.Message)
	utils.Json.DecodeToObj(recvMessage, recvMessageModel)
	if !recvMessageModel.Validate() {
		// 不合法的数据。当没有找到匹配的路由处理
		return nil
	}

	// 判断路由是否存在
	routeArr := strings.Split(recvMessageModel.Route, "/")
	strLen := len(routeArr)
	if strLen != 2 {
		return []byte("illegal route. must be like 'handler/action'")
	}
	handlerName := routeArr[0]
	actionName := routeArr[1]
	hasAction := false // 动作是否存在
	if actionDict, has := component.handlerActionDict[handlerName]; has {
		if _, has = actionDict[actionName]; has {
			hasAction = true
		}
	}
	if !hasAction {
		return []byte("route not found!")
	}

	// 判断控制器是否合法（TODO 这里应该方法注册那里判断）
	handlerType := component.handlerDict[handlerName]
	handlerValue := reflect.New(handlerType.Elem())
	websocketHandlerInterface := handlerValue.Interface().(controllers.WebsocketHandlerInterface)
	if websocketHandlerInterface == nil {
		return []byte("controller must be implement controllers.WebsocketHandlerInterface")
	}
	handlerInterface := handlerValue.Interface().(controllers.HandlerInterface)
	if handlerInterface == nil {
		return []byte("controller must be implement controllers.HandlerInterface")
	}

	// 设置消息
	websocketHandlerInterface.SetSession(session)
	websocketHandlerInterface.SetMessage(recvMessageModel)

	// BeforeAction 一般可用于验证数据
	if !handlerInterface.BeforeAction(actionName) {
		return []byte("illegal request")
	}

	// 调用控制器对应的方法
	action := handlerValue.MethodByName(actionName)
	retValues := action.Call([]reflect.Value{})
	if len(retValues) != 1 || retValues[0].Kind() != reflect.String {
		return []byte("only one argument of type string can be returned")
	}
	sendMessage := retValues[0].Interface().(string)

	// AfterAction 一般可用于对返回数据做进一步的处理
	sendMessage = handlerInterface.AfterAction(actionName, sendMessage)

	return []byte(sendMessage)
}

// 执行 自定义新连接连入方法
func (component *WebsocketServer) callOnConnect(session *models.WebsocketSession) {
	if component.onConnectHandler != nil {
		component.onConnectHandler(session)
	}
}

// 执行 自定义接受到消息方法
func (component *WebsocketServer) callOnMessage(session *models.WebsocketSession, message []byte) {
	if component.onMessageHandler != nil {
		component.onMessageHandler(session, message)
	}
}

// 执行 自定义连接关闭的方法
func (component *WebsocketServer) callOnClose(session *models.WebsocketSession) {
	if component.onCloseHandler != nil {
		component.onCloseHandler(session)
	}
}
