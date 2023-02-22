package http

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
)

type BaseSchema struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Response(ctx *Context, msg string, data ...interface{}) {
	res := &BaseSchema{
		Code:    0,
		Message: msg,
	}
	if len(data) > 0 {
		res.Data = data[0]
	}
	ctx.JSON(200, &res)
}

func ResponseCode(ctx *Context, code int, msg string, data ...interface{}) {
	res := &BaseSchema{
		Code:    code,
		Message: msg,
	}
	if len(data) > 0 {
		res.Data = data[0]
	}
	ctx.JSON(200, &res)
}

func ResponseError(ctx *Context, code int, msg string) {
	base := &BaseSchema{
		Code:    code,
		Message: msg,
	}
	ctx.JSON(200, base)
	//if config.GetEnv() == "debug" {
	//
	//	return
	//}
	//Encrypt(ctx, base)
}

func Encrypt(ctx *Context, res interface{}) {
	resBytes, _ := json.Marshal(res)
	randRes := []byte(fmt.Sprint(rand.Int()))
	has := md5.Sum(randRes)
	md5str := fmt.Sprintf("%x", has)
	resStr := base64.StdEncoding.EncodeToString(resBytes) + md5str[:5]
	resStr = base64.StdEncoding.EncodeToString([]byte(resStr))
	ctx.String(200, resStr)
}
