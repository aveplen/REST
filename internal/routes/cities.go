package routes

import (
	"net/http"

	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func RouteCities(parent *mux.Router, logger *logrus.Logger) {
	cities := parent.PathPrefix("/cities").Subrouter()
	{
		citiesCreate := cities.Methods("POST").Subrouter()
		{
			citiesCreate.HandleFunc("", handles.ApiCitiesPost(logger))
		}
		citiesRead := cities.Methods("GET").Subrouter()
		{
			citiesRead.HandleFunc("", handles.ApiCitiesGet(logger))
			citiesRead.HandleFunc("/{id:[0-9]+}", handles.ApiCitiesGetID(logger))
		}
		citiesUpdate := cities.Methods("PATCH").Subrouter()
		{
			citiesUpdate.HandleFunc("", handles.ApiCitiesPatch(logger))
		}
		citiesDelete := cities.Methods(http.MethodDelete).Subrouter()
		{
			citiesDelete.HandleFunc("", handles.ApiCitiesDelete(logger))
		}
	}
}
