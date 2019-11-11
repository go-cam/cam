package controllers

import "github.com/cinling/cin/core/models"

// xorm's console controller
type XormController struct {
	models.BaseController
}

// Generate ORM classes automatically according to the database
func (controller *XormController) Generate() {

}
