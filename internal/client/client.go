package client

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewClient)

type IClient interface {
	GetAllServices()
}


type Client struct {
	con *ConsulCli
}

func NewClient() *Client {
	return &Client{con: NewConsulCli()}
}
