package utils

import (
	"github.com/gin-gonic/gin"
	"t-blog-back/models"
)

func GetStringFromJson(key string) {

}

func GetLoginUserInfo(c *gin.Context) (user models.User, e error){
	u := c.GetString("user-info")

	jsonP := GetJsonParser()
	e = jsonP.UnmarshalFromString(u, &user)
	return
}