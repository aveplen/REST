package routes

import (
	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteCridentials(parent *mux.Router, s handles.IServer) {
	cridentials := parent.PathPrefix("/cridentials").Subrouter()
	{
		cridentialsCreate := cridentials.Methods("POST").Subrouter()
		{
			cridentialsCreate.HandleFunc("", handles.ApiCridentialsPost(s))
		}
		cridentialsRead := cridentials.Methods("GET").Subrouter()
		{
			cridentialsRead.HandleFunc("", handles.ApiCridentialsGet(s))
			cridentialsRead.HandleFunc("/{id:[0-9]+}", handles.ApiCridentialsGetID(s))
		}
		cridentialsUpdate := cridentials.Methods("PATCH").Subrouter()
		{
			cridentialsUpdate.HandleFunc("", handles.ApiCridentialsPatch(s))
		}
		cridentialsDelete := cridentials.Methods("DELETE").Subrouter()
		{
			cridentialsDelete.HandleFunc("", handles.ApiCridentialsDelete(s))
			cridentialsDelete.HandleFunc("/{id:[0-9]+}", handles.ApiCridentialsDeleteID(s))
		}
	}
}
