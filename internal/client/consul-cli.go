package client

import (
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
)

type ConsulCli struct {
	cli *api.Client
}

func NewConsulCli() *ConsulCli {
	client, err := api.NewClient(&api.Config{
		Address:    "127.0.0.1:8500",
		Scheme:    "http",
		Transport: cleanhttp.DefaultPooledTransport(),
	})
	if err != nil {
		panic(err)
	}
	return &ConsulCli{cli: client}
}

func (c *ConsulCli) GetServiceAll() {
	c.cli.Catalog().Services(&api.QueryOptions{})
}
