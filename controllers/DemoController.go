package controllers

import "github.com/astaxie/beego"

type DemoController struct {
	beego.Controller
}

func (ctrl *DemoController) Get() {
	ctrl.Ctx.WriteString("Hello World!")
}
