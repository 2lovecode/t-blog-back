package e

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const RequestKey = "TankBlogBackRequestKey"

// Body 响应体
type Body struct {
	Code      RCode       `json:"code" example:"200"`
	Msg       string      `json:"msg" example:"成功"`
	Data      interface{} `json:"data"`
	RequestID string      `json:"request_id" example:"buckkbuvvhfijcnmsts0"`
}

// StdError 标准错误
var StdError *TStdError

// TStdError 标准错误
type TStdError struct {
	errorMsg  string
	errorCode RCode
}

// NewStdError 标准错误
func NewStdError(code RCode, err string) *TStdError {
	message := GetMsg(code)
	if message == "" {
		message = err
	}
	return &TStdError{
		errorCode: code,
		errorMsg:  err,
	}
}

func (err *TStdError) Error() string {
	return err.errorMsg
}

// Code 错误码
func (err *TStdError) Code() RCode {
	return err.errorCode
}

// SuccessJSON 成功
func SuccessJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Body{
		Code:      Success,
		Msg:       GetMsg(Success),
		Data:      data,
		RequestID: c.GetString(RequestKey),
	})
}

// FailureJSON 失败
func FailureJSON(c *gin.Context, err error) {
	if errors.As(err, &StdError) {
		stdErr := err.(*TStdError)
		code := stdErr.Code()
		c.JSON(http.StatusBadRequest, Body{
			Code:      code,
			Msg:       stdErr.Error(),
			Data:      nil,
			RequestID: c.GetString(RequestKey),
		})
	} else {
		c.JSON(http.StatusInternalServerError, Body{
			Code:      Error,
			Msg:       GetMsg(Error),
			Data:      nil,
			RequestID: c.GetString(RequestKey),
		})
	}
}
