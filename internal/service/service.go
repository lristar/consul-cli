package service

import (
	"consul-cli/api/consulCli"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewConsulCliService,wire.Bind(new(api.IConsulCliHTTPServer),new(ConsulCliService)),)


