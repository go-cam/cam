package camWebsocket

import (
	"github.com/go-cam/cam/base/camBase"
	"github.com/go-cam/cam/base/camUtils"
	"github.com/go-cam/cam/component"
	"github.com/go-cam/cam/plugin/camPluginContext"
	"github.com/go-cam/cam/plugin/camPluginRouter"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"strings"
)

type WebsocketComponent struct {
	component.Component
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
func (comp *WebsocketComponent) Init(configI camBase.ComponentConfigInterface) {
	comp.Component.Init(configI)

	var ok bool
	comp.config, ok = configI.(*WebsocketComponentConfig)
	if !ok {
		camBase.App.Fatal("WebsocketComponent", "invalid config")
		return
	}
	comp.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	comp.messageParseHandler = comp.defaultMessageParseHandler
	comp.RouterPlugin.Init(&comp.config.RouterPluginConfig)
	comp.ContextPlugin.Init(&comp.config.ContextPluginConfig)
}

// start
func (comp *WebsocketComponent) Start() {
	comp.Component.Start()

	if !comp.config.SslOnly {
		camBase.App.Trace("WebsocketComponent", "listen ws://:"+strconv.FormatUint(uint64(comp.config.Port), 10))
		go comp.listenAndServe()
	}
	if comp.config.IsSslOn {
		camBase.App.Trace("WebsocketComponent", "listen wss://:"+strconv.FormatUint(uint64(comp.config.SslPort), 10))
		go comp.listenAndServeTLS()
	}
}

// new connection
func (comp *WebsocketComponent) handlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := comp.upgrader.Upgrade(w, r, nil)
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
			sendMessage := comp.getSendMessage(session, recvMessage)
			if sendMessage != nil {
				err = conn.WriteMessage(websocket.TextMessage, sendMessage)
				if err != nil {
				}
			}
		}
	}
}

// call controller's action
func (comp *WebsocketComponent) getSendMessage(session camBase.SessionInterface, recvMessage []byte) []byte {
	defer func() {
		if rec := recover(); rec != nil {
			comp.Recover(rec)
		}
	}()

	controllerName, actionName, values := comp.messageParseHandler(recvMessage)

	route := camUtils.Url.HumpToUrl(controllerName) + "/" + camUtils.Url.HumpToUrl(actionName)

	handler := comp.getCustomHandler(route)
	if handler != nil {
		websocketSession, ok := session.(*WebsocketSession)
		if !ok {
			panic("session cannot convert to *WebsocketSession")
		}
		return handler(websocketSession.GetConn())
	}

	controller, action := comp.GetControllerAction(route)
	if controller == nil || action == nil {
		camBase.App.Warn("WebsocketComponent", "404. not found route: "+route)
		return nil
	}

	context := comp.NewContext()

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
func (comp *WebsocketComponent) defaultMessageParseHandler(message []byte) (controllerName string, actionName string, values map[string]interface{}) {
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
func (comp *WebsocketComponent) listenAndServe() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", comp.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(comp.config.Port), 10),
		Handler: mux,
	}
	err := server.ListenAndServe()
	camUtils.Error.Panic(err)
}

// enable server with SSl
func (comp *WebsocketComponent) listenAndServeTLS() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", comp.handlerFunc)
	server := &http.Server{
		Addr:    ":" + strconv.FormatUint(uint64(comp.config.SslPort), 10),
		Handler: mux,
	}
	err := server.ListenAndServeTLS(comp.config.SslCertFile, comp.config.SslKeyFile)
	camUtils.Error.Panic(err)
}

// get custom route handler
func (comp *WebsocketComponent) getCustomHandler(route string) camBase.WebsocketRouteHandler {
	handler, has := comp.config.routeHandlerDict[route]
	if !has {
		return nil
	}
	return handler
}
