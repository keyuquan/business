package utils

import (
	"github.com/rs/xid"
)

var (
	Uuid = xid.New()
)

func NewUuid() string {
	return xid.New().String()
}
