package e

var MsgFlags = map[RCode]string {
	Success : "ok",
	Error : "fail",
	ErrorInvalidParams : "invalid params",
	ErrorExistModule : "module exist",
}
func GetMsg(code RCode) string {
	msg, ok := MsgFlags[code]
	if ok {
		return  msg
	}
	return MsgFlags[Error]
}