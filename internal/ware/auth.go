package ware

import (
	"business/pkg/http"
)

const (
	TokenKey = "merchant_admin:session:%d"
)

func Auth(ctx *http.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		http.ResponseError(ctx, -9998, "没有登录令牌，非法访问")
		ctx.Abort()
		return
	}

}
