package main

import (
	"context"
	"cutlink/database"
	"cutlink/http/repository"
	"cutlink/http/server"
	"cutlink/http/transport"
)

func main() {
	ctx := context.Background()
	conn := database.CreateConnection(ctx)
	repo := repository.NewRepository(ctx, conn)

	handlerList := transport.NewHandlerList(repo)

	httpServer := server.NewHTTPServer(handlerList)

	if err := httpServer.StartServer(); err != nil {
		panic(err)
	}
}
