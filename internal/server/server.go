package server

import (
	"net/http"

	"github.com/aveplen/REST/internal/config"
	"github.com/aveplen/REST/internal/handler"
	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	bindAddr string
}

func NewServer(handler *handler.Handler, cfg config.Server) *Server {
	return &Server{
		router:   handler.InitRoutes(),
		bindAddr: cfg.BindAddr,
	}
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.bindAddr, s.router)
}
