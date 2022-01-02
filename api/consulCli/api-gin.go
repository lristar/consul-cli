package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Srv IConsulCliHTTPServer

type IConsulCliHTTPServer interface {
	Hello(context.Context) (*HelloResp, error)
}

func RegisterConsulCliHTTPServer(g *gin.Engine, svc IConsulCliHTTPServer) *gin.Engine {
	Srv = svc
	r := g.Group("/v1")
	{
		r.GET("/hello", hello)
	}
	return g
}

func hello(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := Srv.Hello(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "失败"})
		return
	}
	c.JSON(http.StatusOK, resp)

}
