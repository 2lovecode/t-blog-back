package controllers

import "github.com/astaxie/beego"

type CategoryController struct {
	beego.Controller
}

func (ctrl *CategoryController) Get() {
	ctrl.Ctx.WriteString("this is category get")
}