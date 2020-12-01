package utils

import (
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
)

// GenUniqueID 生成唯一ID
func GenUniqueID() string {
	id := xid.New()
	return id.String()
}

// GenToken 生成token
func GenToken() string {
	return uuid.NewV4().String()
}
