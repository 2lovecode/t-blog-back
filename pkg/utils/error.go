package utils

import "github.com/astaxie/beego/validation"

func GetFirstErrorMessage(valid validation.Validation) string {
	errors := ""
	if valid.HasErrors() {
		errors = valid.Errors[0].Message
	}
	return errors
}
