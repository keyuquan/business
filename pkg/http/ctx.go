package http

import (
	"business/internal/model"
	"business/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"strings"
)

type Context struct {
	*gin.Context
	RequestID string
}

type Handler func(ctx *Context)

func (c Context) GetClientIP() string {
	forwardHeader := c.Request.Header.Get("x-forwarded-for")
	firstAddress := strings.Split(forwardHeader, ",")[0]
	if net.ParseIP(strings.TrimSpace(firstAddress)) != nil {
		return firstAddress
	}
	return c.ClientIP()
}

func (c *Context) SetAdminInfo(admin *model.TAdmin) {
	c.Set("admin_info", admin)
}

func (c *Context) GetAdminInfo() (*model.TAdmin, bool) {
	adminBytesInterface, ok := c.Get("admin_info")
	if !ok {
		return nil, false
	}
	res, ok := adminBytesInterface.(*model.TAdmin)
	if !ok {
		return nil, false
	}
	return res, ok
}

func (c *Context) Info(msg string) {
	log.Logger.With(zap.String("trace", c.RequestID)).Info(msg)
}

func (c *Context) Infof(msg string, f ...interface{}) {
	log.Logger.With(zap.String("trace", c.RequestID)).Info(fmt.Sprintf(msg, f...))
}

func (c *Context) Error(msg string) {
	log.Logger.With(zap.String("trace", c.RequestID)).Error(msg)
}

func (c *Context) Errorf(msg string, f ...interface{}) {
	log.Logger.With(zap.String("trace", c.RequestID)).Error(fmt.Sprintf(msg, f...))
}

func (c *Context) Fatalf(msg string, f ...interface{}) {
	log.Logger.With(zap.String("trace", c.RequestID)).Fatal(fmt.Sprintf(msg, f...))
}

func (c *Context) GetTraceID() string {
	return c.RequestID
}
