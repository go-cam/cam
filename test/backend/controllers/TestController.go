package controllers

import (
	"github.com/go-cam/cam"
)

// text controller
type TestController struct {
	cam.BaseController
}

func (controller *TestController) Init() {
	controller.BaseController.Init()

	controller.DefaultAction = "Test"
}

// test action
func (controller *TestController) Test() {
	_ = cam.App.Info("title", "content")
	controller.GetApp().Stop()
	controller.Write([]byte("done"))
}

// private func
func (controller *TestController) privateFunc() {

}
