package consul

import (
	"consul-cli/internal/client"
	"consul-cli/internal/server"
	"consul-cli/internal/service"
	"github.com/gin-gonic/gin"
)

var (
	clientClient     *client.Client
	consulCliService service.ConsulCliService
)

type ConBuilder struct {
}

func (c *ConBuilder) NewService() {
	consulCliService = service.NewConsulCliService(clientClient)
}

func (c *ConBuilder) NewHttpServer() *gin.Engine {
	return server.NewHTTPServer(consulCliService)
}

func (c *ConBuilder) NewClient() {
	clientClient = client.NewClient()
}

func (c *ConBuilder) NewDataServer() {

}
