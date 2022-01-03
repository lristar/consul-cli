package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Srv IConsulCliHTTPServer

type IConsulCliHTTPServer interface {
	GetServiceAll(context.Context) (*HelloResp, error)
}

// 过滤器
func HandlerFilter(g *gin.Context) {
	userId := g.Request.Header.Get("user_id")
	if userId == "" {
		g.JSON(http.StatusForbidden, gin.H{"msg": "权限问题"})
		g.Abort()
	}
}

func RegisterConsulCliHTTPServer(g *gin.Engine, svc IConsulCliHTTPServer) *gin.Engine {
	Srv = svc
	r := g.Group("/v1")
	{
		r.GET("/hello", hello)
	}
	e := g.Group("/v2")
	e.Use(HandlerFilter)
	{
		e.GET("hello", hello)
	}

	return g
}

func hello(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := Srv.GetServiceAll(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "失败"})
		return
	}
	c.JSON(http.StatusOK, resp)

}
