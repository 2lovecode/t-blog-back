package e

var MsgFlags = map[RCode]string {
	Success : "ok",
	Error : "fail",
	ErrorInvalidParams : "无效的参数",
	ErrorExistModule : "module exist",
	ErrorInvalidToken: "无效的token",
}
func GetMsg(code RCode) string {
	msg, ok := MsgFlags[code]
	if ok {
		return  msg
	}
	return MsgFlags[Error]
}