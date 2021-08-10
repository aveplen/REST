package handler

import (
	"github.com/aveplen/REST/internal/service"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service    *service.Service
	logger     *logrus.Logger
	middleware map[string]mux.MiddlewareFunc
}

func NewHandler(service *service.Service, logger *logrus.Logger) *Handler {
	return &Handler{
		service:    service,
		logger:     logger,
		middleware: make(map[string]mux.MiddlewareFunc),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("").Subrouter()
	{
		h.RouteAuth(api)
		h.RouteTable(api)
		h.RouteUsers(api)
		h.RoutePersonal(api)
	}
	return api
}

func (h *Handler) RouteAuth(parent *mux.Router) {
	auth := parent.PathPrefix("/auth").Subrouter()
	{
		authSubrouter := auth.Methods("POST").Subrouter()
		{
			authSubrouter.Handle("/login", h.login())
			authSubrouter.Handle("/register", h.register())
		}
	}
}

func (h *Handler) RouteTable(parent *mux.Router) {
	table := parent.PathPrefix("/table").Subrouter()
	{
		tableSubrouter := table.Methods("GET").Subrouter()
		{
			tableSubrouter.Use(h.jwtAuth)
			tableSubrouter.Handle("", h.tableFull())
			tableSubrouter.Handle("/{size:[0-9]+}/{page:[0-9]+}", h.tablePage())
		}
	}
}

func (h *Handler) RouteUsers(parent *mux.Router) {
	users := parent.PathPrefix("/users/{user_id:[0-9]+}").Subrouter()
	users.Use(h.jwtAuth)
	{
		usersGet := users.Methods("GET").Subrouter()
		{
			usersGet.Use(h.adminOrUser)
			usersGet.Handle("", h.userGet())
		}
		usersUpdate := users.Methods("PUT").Subrouter()
		{
			usersUpdate.Use(h.adminOrUser)
			usersUpdate.Handle("", h.userUpdate())
		}
		usersDelete := users.Methods("DELETE").Subrouter()
		{
			usersDelete.Use(h.adminOnly)
			usersDelete.Handle("", h.userDelete())
		}
	}
}

func (h *Handler) RoutePersonal(parent *mux.Router) {
	students := parent.PathPrefix("/users/{user_id:[0-9]+}/personal").Subrouter()
	students.Use(h.jwtAuth)
	{
		studentsCraete := students.Methods("POST").Subrouter()
		{
			studentsCraete.Use(h.adminOrUser)
			studentsCraete.Handle("", h.studentAttach())
		}
		studentsRead := students.Methods("GET").Subrouter()
		{
			studentsRead.Use(h.adminOrUser)
			studentsRead.Handle("", h.studentGet())
		}
		studentsUpdate := students.Methods("PUT").Subrouter()
		{
			studentsUpdate.Use(h.adminOrUser)
			studentsUpdate.Handle("", h.studentUpdate())
		}
		studentsDelete := students.Methods("DELETE").Subrouter()
		{
			studentsDelete.Use(h.adminOrUser)
			studentsDelete.Handle("", h.studentDetach())
		}
	}
}
