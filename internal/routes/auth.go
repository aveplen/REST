package routes

import (
	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteAuth(parent *mux.Router, s handles.IServer) {
	auth := parent.PathPrefix("/auth").Subrouter()
	{
		authCreate := auth.Methods("POST").Subrouter()
		{
			authCreate.HandleFunc("/login", handles.ApiAuthLogin(s))
			authCreate.HandleFunc("/register", handles.ApiAuthRegister(s))
		}
	}
}
