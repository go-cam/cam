package camWebsocket

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/plugin/camPluginContext"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type WebsocketComponent struct {
	camBase.Component
	camPluginRouter.RouterPlugin
	camPluginContext.ContextPlugin

	config *WebsocketComponentConfig

	upgrader websocket.Upgrader // struct of http upgrade to websocket

	// message parse handler
	//
	//
	// message: client send bytes
	//
	// controllerName:  controller name
	// actionName: 		action name
	// values: 			send data, just like post form data
	messageParseHandler camBase.WebsocketMessageParseHandler
}

// init
func (component *WebsocketComponent) Init(configInterface camBase.ComponentConfigInterface) {
	component.Component.Init(configInterface)

	configValue := reflect.ValueOf(configInterface)
	var config *WebsocketComponentConfig
	if configValue.Kind() == reflect.Ptr {
		config = configValue.Interface().(*WebsocketComponentConfig)
	} else if configValue.Kind() == reflect.Struct {
		configStruct := configValue.Interface().(WebsocketComponentConfig)
		config = &configStruct
	} else {
		panic("illegal config")
	}

	component.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	component.messageParseHandler = component.defaultMessageParseHandler
	component.config = config
	component.RouterPlugin.Init(&config.RouterPluginConfig)
	component.ContextPlugin.Init(&config.ContextPluginConfig)
}

// start
func (component *WebsocketComponent) Start() {
	component.Component.Start()

	if !component.config.SslOnly {
		camBase.App.Info("WebsocketComponent", "listen ws://:"+strconv.FormatUint(uint64(component.config.Port), 10))
		go component.listenAndServe()
	}
	if component.config.IsSslOn {
		camBase.App.Info("WebsocketComponent", "listen wss://:"+strconv.FormatUint(uint64(component.config.SslPort), 10))
		go component.listenAndServeTLS()
	}
}

// new connection
func (component *WebsocketComponent) handlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := component.upgrader.Upgrade(w, r, nil)
	if conn == nil || err != nil {
		return
	}

	session := NewWebsocketSession(conn)

	defer func() {
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
			// Use controller or custom message handler to get sendMessage
			sendMessage := component.getSendMessage(session, recvMessage)
			if sendMessage != nil {
				err = conn.WriteMessage(websocket.TextMessage, sendMessage)
				if err != nil {
				}
			}
		}
	}
}

// call controller's action
func (component *WebsocketComponent) getSendMessage(session camBase.SessionInterface, recvMessage []byte) []byte {
	defer func() {
		if rec := recover(); rec != nil {
			panic(rec)
		}
	}()

	controllerName, actionName, values := component.messageParseHandler(recvMessage)

	route := camUtils.Url.HumpToUrl(controllerName) + "/" + camUtils.Url.HumpToUrl(actionName)

	handler := component.getCustomHandler(route)
	if handler != nil {
		websocketSession, ok := session.(*WebsocketSession)
		if !ok {
			panic("session cannot convert to *WebsocketSession")
		}
		return handler(websocketSession.GetConn())
	}

	controller, action := component.GetControllerAction(route)
	if controller == nil || action == nil {
		panic("404")
	}

	context := component.NewContext()

	// init controller
	controller.Init()
	controller.SetContext(context)
	controller.SetSession(session)
	controller.SetValues(values)

	// call before action
	if !controller.BeforeAction(action) {
		return []byte("illegal request")
	}
	action.Call()
	response := controller.AfterAction(action, controller.GetResponse())

	return response
}

// default router parser.
// Parse the received data to: controllerName、actionName、values
func (component *WebsocketComponent) defaultMessageParseHandler(message []byte) (controllerName string, actionName string, values map[string]interface{}) {
	messageModel := new(Message)
	responseModel := new(Response)
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
func (component *WebsocketComponent) listenAndServe() {
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
func (component *WebsocketComponent) listenAndServeTLS() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", component.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(component.config.SslPort), 10),
		Handler: mux,
	}
	err := server.ListenAndServeTLS(component.config.SslCertFile, component.config.SslKeyFile)
	camUtils.Error.Panic(err)
}

// get custom route handler
func (component *WebsocketComponent) getCustomHandler(route string) camBase.WebsocketRouteHandler {
	handler, has := component.config.routeHandlerDict[route]
	if !has {
		return nil
	}
	return handler
}
