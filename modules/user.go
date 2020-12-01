package modules

import (
	"net/http"
	"t-blog-back/logic/user"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserInfoResp 用户数据
type UserInfoResp struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// Login 登录
func Login(c *gin.Context) {
	loginReq := user.LoginForm{}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		utils.FailureJSON(c, err)
		return
	}

	resp, err := user.Login(c, loginReq)

	if err == nil {
		utils.SuccessJSON(c, resp)
	} else {
		utils.FailureJSON(c, err)
	}
}

// RefreshToken 重置token
func RefreshToken(c *gin.Context) {

}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	data := UserInfoResp{
		UID:    "abcd",
		Name:   "2lovecode",
		Avatar: "/assets/images/avatar/avatar-3.jpg",
	}
	code := e.Success

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func NoLoginToken(c *gin.Context) {

}
