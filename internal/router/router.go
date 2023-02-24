package router

import (
	"business/internal/handler"
	"business/internal/ware"
	"business/pkg/http"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			// c.AbortWithStatus(httpStd.StatusNoContent)
			return
		}
		// 处理请求
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	v1 := router.Group("business/v1")
	{
		Get(v1, "/test", handler.Test)
		// 图片上传
		Post(v1, "/upload", handler.Upload)
		// 图片下载
		Get(v1, "/download", handler.Download)

	}
	return router
}

func Get(engine *gin.RouterGroup, path string, handler ...http.Handler) {
	engine.GET(path, ware.Convert(handler...))
}

func Post(engine *gin.RouterGroup, path string, handler ...http.Handler) {
	engine.POST(path, ware.Convert(handler...))
}

func Delete(engine *gin.RouterGroup, path string, handler ...http.Handler) {
	engine.DELETE(path, ware.Convert(handler...))
}
