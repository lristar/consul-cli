package service

import (
	"consul-cli/api/consulCli"
	"consul-cli/internal/client"
	"context"
)

// GreeterService is a greeter service.
type ConsulCliService struct {
	cli *client.Client
}

func NewConsulCliService(cli *client.Client) ConsulCliService {
	return ConsulCliService{cli: cli}
}

func (c *ConsulCliService) GetServiceAll(context.Context) (*api.HelloResp, error) {
	//c.cli.GetServiceAll()
	return &api.HelloResp{Message:"hello"},nil
}
