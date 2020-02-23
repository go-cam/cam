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

	upgrader             websocket.Upgrader         // struct of http upgrade to websocket
	controllerDict       map[string]reflect.Type    // controller reflect.Type dict
	controllerActionDict map[string]map[string]bool // map[controllerName]map[actionName]

	// Deprecated:
	// on new client connection
	onConnectHandler func(conn camBase.ContextInterface)
	// Deprecated:
	// on receive client message
	onMessageHandler func(conn camBase.ContextInterface, recvMessage []byte)
	// Deprecated:
	// on client connection close
	onCloseHandler func(conn camBase.ContextInterface)

	// message parse handler
	//
	//
	// message: client send bytes
	//
	// controllerName:  controller name
	// actionName: 		action name
	// values: 			send data, just like post form data
	messageParseHandler func(message []byte) (controllerName string, actionName string, values map[string]interface{})
}

// init
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

// start
func (component *WebsocketServer) Start() {
	component.Base.Start()

	if !component.config.IsSslOnly {
		go component.listenAndServe()
	}
	if component.config.IsSslOn {
		go component.listenAndServeTLS()
	}
}

//// 设置 接受新连接的方法
//func (component *WebsocketServer) OnConnect(handler func(conn camBase.ContextInterface)) {
//	component.onConnectHandler = handler
//}
//
//// 设置 接受消息的方法
//func (component *WebsocketServer) OnMessage(handler func(conn camBase.ContextInterface, recvMessage []byte)) {
//	component.onMessageHandler = handler
//}
//
//// 设置 关闭连接的方法
//func (component *WebsocketServer) OnClose(handler func(conn camBase.ContextInterface)) {
//	component.onCloseHandler = handler
//}

// new connection
func (component *WebsocketServer) handlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := component.upgrader.Upgrade(w, r, nil)
	if conn == nil || err != nil {
		return
	}

	session := camModels.NewWebsocketSession(conn)
	context := component.config.NewContext()
	context.SetSession(session)
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
		if messageType == websocket.TextMessage || messageType == websocket.BinaryMessage {
			component.callOnMessage(context, recvMessage)
			// Use controller or custom message handler to get sendMessage
			sendMessage := component.getSendMessageByHandler(context, recvMessage)
			if sendMessage != nil {
				err = conn.WriteMessage(websocket.TextMessage, sendMessage)
				if err != nil {
				}
			}
		}
	}
}

// Use controller or custom message handler to get sendMessage
func (component *WebsocketServer) getSendMessageByHandler(context camBase.ContextInterface, recvMessage []byte) (sendMessage []byte) {
	defer func() {
		if rec := recover(); rec != nil {
			sendMessage = camUtils.Json.Encode(rec)
		}
	}()

	// custom message handler
	component.callOnMessage(context, recvMessage)
	if sendMessage != nil {
		return sendMessage
	}

	// call controller's action
	sendMessage = component.callControllerAction(context, recvMessage)

	return sendMessage
}

// call controller's action
func (component *WebsocketServer) callControllerAction(context camBase.ContextInterface, recvMessage []byte) []byte {
	//recvMessageModel := new(camModels.MessageModel)
	//camUtils.Json.DecodeToObj(recvMessage, recvMessageModel)
	//if !recvMessageModel.Validate() {
	//	// 不合法的数据。当没有找到匹配的路由处理
	//	return nil
	//}
	//
	//// 判断路由是否存在
	//routeArr := strings.Split(recvMessageModel.Route, "/")
	//strLen := len(routeArr)
	//if strLen != 2 {
	//	return []byte("illegal route. must be like 'controller/action'")
	//}
	controllerName, actionName, values := component.messageParseHandler(recvMessage)
	hasAction := false
	if actionDict, has := component.controllerActionDict[controllerName]; has {
		if _, has = actionDict[actionName]; has {
			hasAction = true
		}
	}
	if !hasAction {
		return []byte("route not found!")
	}

	// check controller
	controllerType := component.controllerDict[controllerName]
	controllerValue := reflect.New(controllerType.Elem())
	controllerInterface := controllerValue.Interface().(camBase.ControllerInterface)

	// init controller
	controllerInterface.Init()
	controllerInterface.SetApp(component.app)
	controllerInterface.SetContext(context)
	controllerInterface.SetValues(values)

	// call before action
	if !controllerInterface.BeforeAction(actionName) {
		return []byte("illegal request")
	}

	// call action
	action := controllerValue.MethodByName(actionName)
	_ = action.Call([]reflect.Value{})
	response := controllerInterface.Read()

	// call after action
	response = controllerInterface.AfterAction(actionName, response)

	return response
}

// Deprecated:
func (component *WebsocketServer) callOnConnect(context camBase.ContextInterface) {
	if component.onConnectHandler != nil {
		component.onConnectHandler(context)
	}
}

// Deprecated:
func (component *WebsocketServer) callOnMessage(context camBase.ContextInterface, message []byte) {
	if component.onMessageHandler != nil {
		component.onMessageHandler(context, message)
	}
}

// Deprecated:
func (component *WebsocketServer) callOnClose(context camBase.ContextInterface) {
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

// enable server
func (component *WebsocketServer) listenAndServe() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.config.Port), 10),
		Handler: mux,
	}
	err := server.ListenAndServe()
	camUtils.Error.Panic(err)
}

// enable server with SSl
func (component *WebsocketServer) listenAndServeTLS() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.config.SslPort), 10),
		Handler: mux,
	}
	err := server.ListenAndServeTLS(component.config.SslCertFile, component.config.SslKeyFile)
	camUtils.Error.Panic(err)
}
