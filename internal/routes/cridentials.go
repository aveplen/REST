package routes

import (
	"net/http"

	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func RouteCridentials(parent *mux.Router, logger *logrus.Logger) {
	cridentials := parent.PathPrefix("/cridentials").Subrouter()
	{
		cridentialsCreate := cridentials.Methods("POST").Subrouter()
		{
			cridentialsCreate.HandleFunc("", handles.ApiCridentialsPost(logger))
		}
		cridentialsRead := cridentials.Methods("GET").Subrouter()
		{
			cridentialsRead.HandleFunc("", handles.ApiCridentialsGet(logger))
			cridentialsRead.HandleFunc("/{id:[0-9]+}", handles.ApiCridentialsGetID(logger))
		}
		cridentialsUpdate := cridentials.Methods("PATCH").Subrouter()
		{
			cridentialsUpdate.HandleFunc("", handles.ApiCridentialsPatch(logger))
		}
		cridentialsDelete := cridentials.Methods(http.MethodDelete).Subrouter()
		{
			cridentialsDelete.HandleFunc("", handles.ApiCridentialsDelete(logger))
		}
	}
}
