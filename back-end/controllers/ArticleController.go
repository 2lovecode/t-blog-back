package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller
}

func (ctrl *ArticleController) Get() {
	ctrl.Ctx.WriteString("this is article get")
}

func (ctrl *ArticleController) Post() {
	ctrl.Ctx.WriteString("this article post")
}

func (ctrl *ArticleController) List() {
	ctrl.Ctx.WriteString("this is article list")
}

