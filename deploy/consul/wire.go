// The build tag makes sure the stub is not built in the final build.
// +build wireinject
//go:generate wire

package consul

import (
	"consul-cli/internal/client"
	"consul-cli/internal/server"
	"consul-cli/internal/service"
	"github.com/google/wire"
)

// initApp init consul application.
func InitApp() (*app) {
	panic(wire.Build(client.ProviderSet,service.ProviderSet,server.ProviderSet,newApp))
}
