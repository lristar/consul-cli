package client

import (
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

var ProviderSet = wire.NewSet(NewClient)

type IClient interface {
	GetAllServices()
}

type Client struct {
	con *api.Client
}

func NewClient() *Client {
	return &Client{con: NewConsulCli()}
}
