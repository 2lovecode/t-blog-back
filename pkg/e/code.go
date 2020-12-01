package e

// RCode code码
type RCode int64

const (
	// Success 成功
	Success RCode = 200000

	// Error 异常
	Error = 500000

	// Failure 错误
	Failure = 400000
	// FailureInvalidParams 无效参数
	FailureInvalidParams = 400100
	// FailureExistModule 模块不存在
	FailureExistModule = 400200
	// FailureInvalidUserOrPass 用户名密码错误
	FailureInvalidUserOrPass = 400300
	// FailureInvalidToken 无效token
	FailureInvalidToken = 400301
)

const (
	// StatusSuccess 成功
	StatusSuccess = "success"
	// StatusFailure 失败
	StatusFailure = "failure"
)

// MsgFlags 错误码和message对应关系
var MsgFlags = map[RCode]string{
	Success:                  "成功",
	Error:                    "失败",
	FailureInvalidParams:     "无效的参数",
	FailureExistModule:       "模块不存在",
	FailureInvalidUserOrPass: "用户名或密码错误",
	FailureInvalidToken:      "无效的token",
}

// GetMsg 通过错误码获取message
func GetMsg(code RCode) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return ""
}
