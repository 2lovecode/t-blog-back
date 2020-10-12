package utils

import "github.com/rs/xid"

func GenUniqueID() string {
	id := xid.New()
	return id.String()
}