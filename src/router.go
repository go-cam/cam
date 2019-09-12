package cin

import "cin/src/controllers"

// 路由器
type router struct {
	handlerList []controllers.HandlerInterface
}

// 新建路由器
func newRouter() *router {
	r := new(router)
	r.handlerList = []controllers.HandlerInterface{}
	return r
}

// 注册控制器
func (r *router) Register(controller controllers.HandlerInterface) {
	r.handlerList = append(r.handlerList, controller)
}