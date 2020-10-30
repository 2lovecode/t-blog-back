package modules

import (
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
	"t-blog-back/pkg/utils"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// LoginReq 请求
type LoginReq struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

// LoginResp 响应
type LoginResp struct {
	Token string `json:"token"`
}

// UserInfoResp 用户数据
type UserInfoResp struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// Login 登录
func Login(c *gin.Context) {

	loginReq := LoginReq{}
	if err := c.ShouldBindJSON(&loginReq); err == nil {
		code := e.Success
		user := models.User{}

		data := LoginResp{}
		if u, err := user.FindUserByName(loginReq.UserName); err == nil && bcrypt.CompareHashAndPassword([]byte(u.Pass), []byte(loginReq.PassWord)) == nil {
			token := uuid.NewV4().String()
			_, err = user.UpdateToken(loginReq.UserName, token)
			data.Token = token
		} else {
			code = e.Error
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	} else {
		utils.AbortWithMessage(c, http.StatusBadRequest, e.Error, err.Error(), nil)
	}

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
