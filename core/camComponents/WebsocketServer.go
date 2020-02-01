package camComponents

import (
	"github.com/go-cam/cam/core/camBase"
	"github.com/go-cam/cam/core/camConfigs"
	"github.com/go-cam/cam/core/camModels"
	"github.com/go-cam/cam/core/camUtils"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type WebsocketServer struct {
	Base

	config *camConfigs.WebsocketServer
	port   uint16 // websocket 监听端口

	upgrader             websocket.Upgrader         // websocket http 升级为 websocket 的方法
	controllerDict       map[string]reflect.Type    // 控制器反射map
	controllerActionDict map[string]map[string]bool // 控制器 => 方法 => 是否存在（注册时记录）

	onConnectHandler func(conn *camModels.Context)                     // 自定义方法：有新连接连入
	onMessageHandler func(conn *camModels.Context, recvMessage []byte) // 自定义方法：收到消息
	onCloseHandler   func(conn *camModels.Context)                     // 自定义方法：连接被关闭

	// 传输消息解析器
	// message: 客户端发送过来的消息
	// controllerName: 控制器名字
	// actionName: 控制器方法名字
	// values: 传输的参数
	messageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})
}

// 使用配置 初始化数据
func (component *WebsocketServer) Init(configInterface camBase.ConfigComponentInterface) {
	component.Base.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *camConfigs.WebsocketServer
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*camConfigs.WebsocketServer)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(camConfigs.WebsocketServer)
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
func (component *WebsocketServer) OnConnect(handler func(conn *camModels.Context)) {
	component.onConnectHandler = handler
}

// 设置 接受消息的方法
func (component *WebsocketServer) OnMessage(handler func(conn *camModels.Context, recvMessage []byte)) {
	component.onMessageHandler = handler
}

// 设置 关闭连接的方法
func (component *WebsocketServer) OnClose(handler func(conn *camModels.Context)) {
	component.onCloseHandler = handler
}

// 处理收到消息的方法
func (component *WebsocketServer) handlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := component.upgrader.Upgrade(w, r, nil)
	if conn == nil || err != nil {
		return
	}

	session := camModels.NewWebsocketSession(conn)
	context := camModels.NewContext(session)
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
func (component *WebsocketServer) getSendMessageByHandler(context *camModels.Context, recvMessage []byte) (sendMessage []byte) {
	defer func() {
		if rec := recover(); rec != nil {
			sendMessage = camUtils.Json.Encode(rec)
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
func (component *WebsocketServer) callControllerAction(context *camModels.Context, recvMessage []byte) []byte {
	recvMessageModel := new(camModels.MessageModel)
	camUtils.Json.DecodeToObj(recvMessage, recvMessageModel)
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
	controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)

	// init controller
	controllerInterface.Init()
	controllerInterface.SetApp(component.app)
	controllerInterface.SetContext(context)
	controllerInterface.SetValues(values)

	// BeforeAction 一般可用于验证数据
	if !controllerInterface.BeforeAction(actionName) {
		return []byte("illegal request")
	}

	// 调用控制器对应的方法
	action := controllerValue.MethodByName(actionName)
	_ = action.Call([]reflect.Value{})
	response := controllerInterface.Read()

	// AfterAction 一般可用于对返回数据做进一步的处理
	response = controllerInterface.AfterAction(actionName, response)

	return response
}

// 执行 自定义新连接连入方法
func (component *WebsocketServer) callOnConnect(context *camModels.Context) {
	if component.onConnectHandler != nil {
		component.onConnectHandler(context)
	}
}

// 执行 自定义接受到消息方法
func (component *WebsocketServer) callOnMessage(context *camModels.Context, message []byte) {
	if component.onMessageHandler != nil {
		component.onMessageHandler(context, message)
	}
}

// 执行 自定义连接关闭的方法
func (component *WebsocketServer) callOnClose(context *camModels.Context) {
	if component.onCloseHandler != nil {
		component.onCloseHandler(context)
	}
}

// 默认路由解析手柄
func (component *WebsocketServer) defaultRouteParseHandler(message []byte) (controllerName string, actionName string, values map[string]interface{}) {
	messageModel := new(camModels.MessageModel)
	responseModel := new(camModels.ResponseModel)
	camUtils.Json.DecodeToObj(message, messageModel)
	camUtils.Json.DecodeToObj([]byte(messageModel.Data), responseModel)

	if messageModel.Route == "" {
		return "", "", responseModel.Values
	}
	if !strings.Contains(messageModel.Route, "/") {
		return messageModel.Route, "", responseModel.Values
	}
	tmpArr := strings.Split(messageModel.Route, "/")
	return camUtils.Url.UrlToHump(tmpArr[0]), camUtils.Url.UrlToHump(tmpArr[1]), responseModel.Values
}
