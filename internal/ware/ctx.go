package ware

import (
	"business/pkg/http"
	"business/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Convert(fs ...http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &http.Context{
			Context: c,
		}
		ctx.RequestID = utils.NewUuid()
		maxIndex := len(fs) - 1
		for i := maxIndex; i >= 0; i-- {
			fs[i](ctx)
			if ctx.IsAborted() {
				return
			}
		}
	}
}
