package main

import (
	"consul-cli/deploy"
	"consul-cli/deploy/consul"
)

func main() {
	// Get a new client
	con := consul.ConBuilder{}
	builder := deploy.Builder{&con}
	app := builder.InitApp()
	_ = app.G.Run(":9000")

}
