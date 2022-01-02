package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
)

func main() {
	// Get a new client
	client, err := api.NewClient(&api.Config{
		Address:   "192.168.195.145:8500",
		Scheme:    "http",
		Transport: cleanhttp.DefaultPooledTransport(),
	})
	if err != nil {
		return
	}
	services,meta,err:=client.Catalog().Services(&api.QueryOptions{NodeMeta: map[string]string{"Cluster":"F"}})
	if err!=nil{
		return
	}
	fmt.Println("service:",services)
	fmt.Println("meta:",meta)
}
