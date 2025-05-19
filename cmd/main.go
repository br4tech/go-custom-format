package main

import (
	"github.com/br4tech/go-custom-format/config"
	"github.com/br4tech/go-custom-format/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewechoServer(&cfg).Start()
}
