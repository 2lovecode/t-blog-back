package modules

import (
	"t-blog-back/pkg/utils"

	"github.com/gin-gonic/gin"
)

// AboutAuthor 关于作者信息
type AboutAuthor struct {
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Email   string `json:"email"`
	Intro   string `json:"intro"`
	Article int    `json:"article"`
}

// About 关于作者
func About(ctx *gin.Context) {
	utils.SuccessJSON(ctx, AboutAuthor{
		Name:    "2lovecode",
		Avatar:  "http://www.2lovecode.com:9999/tank-blog/avatar/avatar-1.jpg",
		Email:   "liu546hao@163.com",
		Intro:   "搞点事情～",
		Article: 0,
	})
}
