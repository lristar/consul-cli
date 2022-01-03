package client

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
)

func NewConsulCli() *api.Client {
	client, err := api.NewClient(&api.Config{
		Address:   "192.168.195.145:8500",
		Scheme:    "http",
		Transport: cleanhttp.DefaultPooledTransport(),
	})
	if err != nil {
		panic(err)
	}
	return client
}

func (c *Client) GetServiceAll() {
	data, meta, err := c.con.Catalog().Services(&api.QueryOptions{NodeMeta: map[string]string{"Cluster":"D"}})
	if err != nil {
		return
	}
	fmt.Println("data:", data)
	fmt.Println("meta:", meta)
}
