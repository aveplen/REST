package server

import (
	"net/http"

	"github.com/aveplen/REST/internal/config"
	"github.com/aveplen/REST/internal/routes"
	"github.com/aveplen/REST/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store.NewStore(config.Pg),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(s.config.Log); err != nil {
		return err
	}
	if err := s.store.Open(); err != nil {
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

func (s *Server) GetLogger() *logrus.Logger {
	return s.logger
}

func (s *Server) GetStore() *store.Store {
	return s.store
}

func (s *Server) configureRouter() {
	api := s.router.PathPrefix("/api").Subrouter()
	{
		routes.RouteCities(api, s)
		routes.RouteSchools(api, s)
		routes.RouteStudents(api, s)
		routes.RouteUsers(api, s)
	}
}
