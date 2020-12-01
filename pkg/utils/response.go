package utils

import (
	"errors"
	"net/http"
	"t-blog-back/pkg/e"

	"github.com/gin-gonic/gin"
)

// BodyJSON 返回的json结构体
type BodyJSON struct {
	Status    string      `json:"status" example:"success"`
	Code      e.RCode     `json:"code" example:"200"`
	Message   string      `json:"msg" example:"成功"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id" example:"buckkbuvvhfijcnmsts0"`
}

// SuccessJSON 成功返回
func SuccessJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BodyJSON{
		Status:    e.StatusSuccess,
		Code:      e.Success,
		Message:   "",
		Data:      data,
		RequestID: "",
	})
}

// SuccessJSONWithMessage 成功返回
func SuccessJSONWithMessage(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, BodyJSON{
		Status:    e.StatusSuccess,
		Code:      e.Success,
		Message:   msg,
		Data:      data,
		RequestID: "",
	})
}

// FailureJSON 失败返回
func FailureJSON(c *gin.Context, err error) {
	FailureJSONWithHTTPCode(c, err, http.StatusInternalServerError)
}

// FailureJSONWithHTTPCode 失败返回
func FailureJSONWithHTTPCode(c *gin.Context, err error, httpCode int) {
	if errors.As(err, &e.StdError) {
		stdErr := err.(*e.TStdError)
		code := stdErr.Code()
		msg := stdErr.Error()
		if msg == "" {
			msg = e.GetMsg(code)
		}
		c.JSON(http.StatusOK, BodyJSON{
			Code:      code,
			Status:    e.StatusFailure,
			Message:   msg,
			Data:      nil,
			RequestID: "",
		})
	} else {
		c.JSON(httpCode, BodyJSON{
			Code:      e.Failure,
			Status:    e.StatusFailure,
			Message:   e.GetMsg(e.Failure),
			Data:      nil,
			RequestID: "",
		})
	}
}

// AbortJSON 失败返回
func AbortJSON(c *gin.Context, err error, httpCode int) {
	c.Abort()
	FailureJSONWithHTTPCode(c, err, httpCode)
}

func Success(c *gin.Context, code e.RCode, eMsg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   e.GetMsg(code),
		"error": eMsg,
		"data":  data,
	})
}

func Abort(c *gin.Context, httpCode int, errCode e.RCode, data interface{}) {
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
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
		"msg":  errMsg,
		"data": data,
	})
	return
}
