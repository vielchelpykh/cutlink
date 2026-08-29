package main

import (
	"cutlink/http/server"
	"cutlink/http/transport"
)

func main() {
	allLinks := transport.NewList()
	handlerList := transport.NewHandlerList(allLinks)
	httpServer := server.NewHTTPServer(handlerList)
	if err := httpServer.StartServer(); err != nil {
		panic(err)
	}
}
