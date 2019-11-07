package components

import (
	"github.com/cinling/cin/base"
	"github.com/cinling/cin/configs"
	"github.com/cinling/cin/models"
	"github.com/cinling/cin/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type WebsocketServer struct {
	Base

	config *configs.WebsocketServer
	port   uint16 // websocket 监听端口

	upgrader             websocket.Upgrader         // websocket http 升级为 websocket 的方法
	controllerDict       map[string]reflect.Type    // 控制器反射map
	controllerActionDict map[string]map[string]bool // 控制器 => 方法 => 是否存在（注册时记录）

	onConnectHandler func(conn *models.Context)                     // 自定义方法：有新连接连入
	onMessageHandler func(conn *models.Context, recvMessage []byte) // 自定义方法：收到消息
	onCloseHandler   func(conn *models.Context)                     // 自定义方法：连接被关闭

	// 传输消息解析器
	// message: 客户端发送过来的消息
	// controllerName: 控制器名字
	// actionName: 控制器方法名字
	// values: 传输的参数
	messageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})
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
	component.controllerDict = map[string]reflect.Type{}
	component.controllerActionDict = map[string]map[string]bool{}
	component.onConnectHandler = nil
	component.onMessageHandler = nil
	component.onCloseHandler = nil
	component.messageParseHandler = component.defaultRouteParseHandler
	component.config = config

	// 注册处理器（控制器）
	component.controllerDict, component.controllerActionDict = common.getControllerDict(config.ControllerList)
	component.onMessageHandler = config.OnWebsocketMessageHandler
	if config.MessageParseHandler != nil {
		component.messageParseHandler = config.MessageParseHandler
	}
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

// 设置 接受新连接的方法
func (component *WebsocketServer) OnConnect(handler func(conn *models.Context)) {
	component.onConnectHandler = handler
}

// 设置 接受消息的方法
func (component *WebsocketServer) OnMessage(handler func(conn *models.Context, recvMessage []byte)) {
	component.onMessageHandler = handler
}

// 设置 关闭连接的方法
func (component *WebsocketServer) OnClose(handler func(conn *models.Context)) {
	component.onCloseHandler = handler
}

// 处理收到消息的方法
func (component *WebsocketServer) handlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := component.upgrader.Upgrade(w, r, nil)
	if conn == nil || err != nil {
		return
	}

	session := models.NewWebsocketSession(conn)
	context := models.NewContext(session)
	component.callOnConnect(context)

	defer func() {
		component.callOnClose(context)
		session.Destroy()
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
			component.callOnMessage(context, recvMessage)
			// 处理客户端发送的请求并返回数据。如果返回 nil ，则代表不会返回数据给客户端
			sendMessage := component.getSendMessageByHandler(context, recvMessage)
			if sendMessage != nil {
				err = conn.WriteMessage(websocket.TextMessage, sendMessage)
				if err != nil {
				}
			}
		}
	}
}

// 使用处理器获取返回数据
func (component *WebsocketServer) getSendMessageByHandler(context *models.Context, recvMessage []byte) (sendMessage []byte) {
	defer func() {
		if rec := recover(); rec != nil {
			sendMessage = utils.Json.Encode(rec)
		}
	}()

	// 使用自定义方法处理
	component.callOnMessage(context, recvMessage)
	if sendMessage != nil {
		return sendMessage
	}

	// 如果没有数据返回，则使用处理器尝试匹配路由
	sendMessage = component.callControllerAction(context, recvMessage)

	return sendMessage
}

// 获取返回结果
func (component *WebsocketServer) callControllerAction(context *models.Context, recvMessage []byte) []byte {
	recvMessageModel := new(models.MessageModel)
	utils.Json.DecodeToObj(recvMessage, recvMessageModel)
	if !recvMessageModel.Validate() {
		// 不合法的数据。当没有找到匹配的路由处理
		return nil
	}

	// 判断路由是否存在
	routeArr := strings.Split(recvMessageModel.Route, "/")
	strLen := len(routeArr)
	if strLen != 2 {
		return []byte("illegal route. must be like 'controller/action'")
	}
	controllerName, actionName, values := component.messageParseHandler(recvMessage)
	hasAction := false // 动作是否存在
	if actionDict, has := component.controllerActionDict[controllerName]; has {
		if _, has = actionDict[actionName]; has {
			hasAction = true
		}
	}
	if !hasAction {
		return []byte("route not found!")
	}

	// 判断控制器是否合法
	controllerType := component.controllerDict[controllerName]
	controllerValue := reflect.New(controllerType.Elem())
	controllerInterface := controllerValue.Interface().(base.ControllerInterface)

	controllerInterface.SetContext(context)
	controllerInterface.SetValues(values)

	// BeforeAction 一般可用于验证数据
	if !controllerInterface.BeforeAction(actionName) {
		return []byte("illegal request")
	}

	// 调用控制器对应的方法
	action := controllerValue.MethodByName(utils.Url.HumpToUrl(actionName))
	retValues := action.Call([]reflect.Value{})
	if len(retValues) != 1 || retValues[0].Kind() != reflect.String {
		return []byte("only one argument of type string can be returned")
	}
	sendMessage := retValues[0].Interface().([]byte)

	// AfterAction 一般可用于对返回数据做进一步的处理
	sendMessage = controllerInterface.AfterAction(actionName, sendMessage)

	return sendMessage
}

// 执行 自定义新连接连入方法
func (component *WebsocketServer) callOnConnect(context *models.Context) {
	if component.onConnectHandler != nil {
		component.onConnectHandler(context)
	}
}

// 执行 自定义接受到消息方法
func (component *WebsocketServer) callOnMessage(context *models.Context, message []byte) {
	if component.onMessageHandler != nil {
		component.onMessageHandler(context, message)
	}
}

// 执行 自定义连接关闭的方法
func (component *WebsocketServer) callOnClose(context *models.Context) {
	if component.onCloseHandler != nil {
		component.onCloseHandler(context)
	}
}

// 默认路由解析手柄
func (component *WebsocketServer) defaultRouteParseHandler(message []byte) (controllerName string, actionName string, values map[string]interface{}) {
	messageModel := new(models.MessageModel)
	responseModel := new(models.ResponseModel)
	utils.Json.DecodeToObj(message, messageModel)
	utils.Json.DecodeToObj([]byte(messageModel.Data), responseModel)

	if messageModel.Route == "" {
		return "", "", responseModel.Values
	}
	if !strings.Contains(messageModel.Route, "/") {
		return messageModel.Route, "", responseModel.Values
	}
	tmpArr := strings.Split(messageModel.Route, "/")
	return utils.Url.UrlToHump(tmpArr[0]), utils.Url.UrlToHump(tmpArr[1]), responseModel.Values
}
