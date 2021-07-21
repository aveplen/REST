package server

import (
	"io"
	"net/http"

	"github.com/aveplen/REST/internal/config"
	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(s.config.Log); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting server")

	return http.ListenAndServe(s.config.Srv.BindAddr, s.router)
}

func (s *Server) configureLogger(config config.Logrus) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/hello", handleHello())
	s.router.HandleFunc("/api/students", handles.ApiStudentsGet()).Methods("GET")
	s.router.HandleFunc("/api/students/{id:[0-9]+}", handles.ApiStudentsGetID(s.logger)).Methods("GET")
	s.router.HandleFunc("/api/students", handles.ApiStudentsPost()).Methods("POST")
}

func handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}