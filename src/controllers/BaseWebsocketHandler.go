package controllers

import "models"

// websocket server 处理器接口
type WebsocketHandlerInterface interface {
	SetSession(session *models.WebsocketSession)
	SetMessage(message *models.Message)
}

// websocket 处理器
type BaseWebsocketHandler struct {
	WebsocketHandlerInterface
	BaseHandler

	session               *models.WebsocketSession // websocket session
	message               *models.Message          // 本次接收的的数据封装
	response              *models.Response         // 消息中数据的封装（如果符合格式就有。否则没有）
	messageDataIsResponse bool                     // message 的 data 数据是否是 response 类型
}

// 设置 websocket session
func (handler *BaseWebsocketHandler) SetSession(session *models.WebsocketSession) {
	handler.session = session
}

// 获取 websocket session
func (handler *BaseWebsocketHandler) GetSession(session *models.WebsocketSession) *models.WebsocketSession {
	return handler.session
}

// 设置 客户端消息
func (handler *BaseWebsocketHandler) SetMessage(message *models.Message) {
	handler.message = message
}

// 获取 客户端消息
func (handler *BaseWebsocketHandler) GetMessage() *models.Message {
	return handler.message
}

// 初始化方法
func (handler *BaseWebsocketHandler) Init() {
	handler.BaseHandler.Init()
	handler.session = nil
	handler.message = nil
	handler.response = nil
	handler.messageDataIsResponse = true
}