package controllers

import (
	"fmt"
	"github.com/go-cam/cam"
)

// text controller
type TestController struct {
	cam.BaseController
}

// test action
func (controller *TestController) Test() {
	fmt.Println("test")
	controller.Write([]byte("test"))
}

// private func
func (controller *TestController) privateFunc() {

}
