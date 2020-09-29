package utils

import (
	"github.com/gin-gonic/gin"
	"t-blog-back/pkg/e"
)

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