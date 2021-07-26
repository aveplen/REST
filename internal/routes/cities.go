package routes

import (
	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteCities(parent *mux.Router, s handles.IServer) {
	cities := parent.PathPrefix("/cities").Subrouter()
	{
		citiesCreate := cities.Methods("POST").Subrouter()
		{
			citiesCreate.HandleFunc("", handles.ApiCitiesPost(s))
		}
		citiesRead := cities.Methods("GET").Subrouter()
		{
			citiesRead.HandleFunc("", handles.ApiCitiesGet(s))
			citiesRead.HandleFunc("/{id:[0-9]+}", handles.ApiCitiesGetID(s))
		}
		citiesUpdate := cities.Methods("PATCH").Subrouter()
		{
			citiesUpdate.HandleFunc("", handles.ApiCitiesPatch(s))
		}
		citiesDelete := cities.Methods("DELETE").Subrouter()
		{
			citiesDelete.HandleFunc("", handles.ApiCitiesDelete(s))
			citiesDelete.HandleFunc("/{id:[0-9]+}", handles.ApiCitiesDeleteID(s))
		}
	}
}
