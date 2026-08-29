package server

import (
	"cutlink/http/transport"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HandlerList *transport.HandlerList
}

func NewHTTPServer(handlerList *transport.HandlerList) *HTTPServer {
	return &HTTPServer{
		HandlerList: handlerList,
	}
}

func (server *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/links").Methods("POST").HandlerFunc(server.HandlerList.HandlerCreateShortLink)
	router.Path("/links/{title}").Methods("PATCH").HandlerFunc(server.HandlerList.HandlerFollowLink)
	router.Path("/links/{title}").Methods("GET").HandlerFunc(server.HandlerList.HandlerGetStatistic)

	if err := http.ListenAndServe(":5000", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
