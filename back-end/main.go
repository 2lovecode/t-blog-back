package main

import (
	"github.com/astaxie/beego"
	. "github.com/t-blog-back/back-end/controllers"
)
func main() {
	beego.Router("/", &DemoController{})
	beego.Router("/article", &ArticleController{})

	beego.Router("/article/list", &ArticleController{}, "get:List")
	beego.Router("/category", &CategoryController{})
	beego.Run()
}
