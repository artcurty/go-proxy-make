package main

import (
	_ "github.com/artcurty/go-proxy-make/cmd/api/generated"
	"github.com/artcurty/go-proxy-make/cmd/api/server"
)

func main() {
	server.StartServer()
}
