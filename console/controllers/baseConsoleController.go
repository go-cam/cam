package controllers

import (
	"github.com/cinling/cin/core/components"
	"github.com/cinling/cin/core/models"
)

// 基础控制器
type baseConsoleController struct {
	models.BaseController
}

// get database component instance
func (controller MigrateController) getDatabaseComponent() *components.Database {
	ins := controller.GetAppInterface().GetComponentByName("db")
	if ins == nil {
		return nil
	}
	return ins.(*components.Database)
}
