package main

import "consul-cli/deploy/consul"

func main() {
	// Get a new client
	s := consul.InitApp()
	_ = s.G.Run(":9000")

}
