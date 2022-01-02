package server

import (
	"consul-cli/api/consulCli"
	"consul-cli/internal/service"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(svc service.ConsulCliService) *gin.Engine {
	e := gin.Default()
	return api.RegisterConsulCliHTTPServer(e,&svc)
}
