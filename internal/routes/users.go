package routes

import (
	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteUsers(parent *mux.Router, s handles.IServer) {
	users := parent.PathPrefix("/Users").Subrouter()
	{
		usersCreate := users.Methods("POST").Subrouter()
		{
			usersCreate.HandleFunc("", handles.ApiUsersPost(s))
		}
		usersDelete := users.Methods("DELETE").Subrouter()
		{
			usersDelete.HandleFunc("/{id:[0-9]+}", handles.ApiUsersDelete(s))
		}
	}
}
