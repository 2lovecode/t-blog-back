package modules

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"t-blog-back/models"
	"t-blog-back/pkg/e"
)

type LoginResp struct {
	Token string `json:"token"`
}

type UserInfoResp struct {
	UId string `json:"uid"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

func Login(c *gin.Context) {

	userName := c.PostForm("user")
	userPass := c.PostForm("pass")
	code := e.Success

	user := models.User{}

	data := LoginResp{}
	if u, err := user.FindUserByName(userName); err == nil && bcrypt.CompareHashAndPassword([]byte(u.Pass), []byte(userPass)) == nil {
		token := uuid.NewV4().String()
		_, err = user.UpdateToken(userName, token)
		data.Token = token
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func UserInfo(c *gin.Context) {
	data := UserInfoResp{
		UId: "abcd",
		Name: "2lovecode",
		Avatar: "/assets/images/avatar/avatar-3.jpg",
	}
	code := e.Success

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
