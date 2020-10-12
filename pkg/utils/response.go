package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t-blog-back/pkg/e"
)

func Success(c *gin.Context, code e.RCode, eMsg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"error": eMsg,
		"data" : data,
	})
}

func Abort(c *gin.Context, httpCode int, errCode e.RCode, data interface{}) {
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code": errCode,
		"msg":	e.GetMsg(errCode),
		"data": data,
	})
	return
}

func AbortWithMessage(c *gin.Context, httpCode int, errCode e.RCode, errMsg string, data interface{}) {
	if errMsg == "" {
		errMsg = e.GetMsg(errCode)
	}
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code": errCode,
		"msg":	errMsg,
		"data": data,
	})
	return
}