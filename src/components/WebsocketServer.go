package components

import (
	"cin/src/alias"
	"cin/src/configs"
	"cin/src/controllers"
	"cin/src/models"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type WebsocketServer struct {
	port uint16                    // websocket 监听端口

	upgrader    websocket.Upgrader      // websocket http 升级为 websocket 的方法
	handlerDict map[string]reflect.Type // 控制器反射map

	onConnectHandler func(conn *models.WebsocketSession)                     // 自定义方法：有新连接连入
	onMessageHandler func(conn *models.WebsocketSession, recvMessage []byte) // 自定义方法：收到消息
	onCloseHandler   func(conn *models.WebsocketSession)                     // 自定义方法：连接被关闭

	mode alias.WebsocketServerMode // 运行模式
}

// 使用配置初始化数据
func NewWebsocketServer(config *configs.WebsocketServer) *WebsocketServer {
	component := new(WebsocketServer)
	component.port = config.Port
	component.mode = config.Mode
	component.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	component.handlerDict = map[string]reflect.Type{}
	component.onConnectHandler = nil
	component.onMessageHandler = nil
	component.onCloseHandler = nil
	return component
}

// 注册处理器。注册后将根据路由自动匹配调用方法
func (component *WebsocketServer) Register(handlerInterface controllers.WebsocketHandlerInterface) {
	t := reflect.TypeOf(handlerInterface)
	handlerType := t.Elem() // 获取实体
	handlerName := handlerType.Name()
	handlerName = strings.TrimSuffix(handlerName, "Handler")
	handlerName = strings.TrimSuffix(handlerName, "Controller")
	handlerName = strings.ToLower(handlerName)
	component.handlerDict[handlerName] = t
}

// 设置 接受消息的方法
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

// 开启 websocket 监听
func (component *WebsocketServer) Run() {
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
		_, recvMessage, err = conn.ReadMessage()
		if err != nil {
			break
		}
		component.callOnMessage(session, recvMessage)

		// 处理客户端发送的请求并返回数据。如果返回 nil ，则代表不会返回数据给客户端
		sendMessage := component.getSendMessageByHandler(conn, recvMessage)
		if sendMessage != nil {
			err = conn.WriteMessage(websocket.TextMessage, sendMessage)
			if err != nil {
			}
		}
	}
}

// 使用处理器获取返回数据
func (component *WebsocketServer) getSendMessageByHandler(conn *websocket.Conn, recvMessage []byte) (sendMessage []byte) {
	defer func() {
		if rec := recover(); rec != nil {
			sendMessage, _ = json.Marshal(rec)
		}
	}()
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