package modules

import (
	"t-blog-back/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserInfoResp 用户数据
type UserInfoResp struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	data := UserInfoResp{
		UID:    "abcd",
		Name:   "2lovecode",
		Avatar: "/assets/images/avatar/avatar-3.jpg",
	}
	utils.SuccessJSON(c, data)
}
