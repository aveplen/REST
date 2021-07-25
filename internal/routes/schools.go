package routes

import (
	"net/http"

	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteSchools(parent *mux.Router, s handles.IServer) {
	schools := parent.PathPrefix("/schools").Subrouter()
	{
		schoolsCreate := schools.Methods("POST").Subrouter()
		{
			schoolsCreate.HandleFunc("", handles.ApiSchoolsPost(s))
		}
		schoolsRead := schools.Methods("GET").Subrouter()
		{
			schoolsRead.HandleFunc("", handles.ApiSchoolsGet(s))
			schoolsRead.HandleFunc("/{id:[0-9]+}", handles.ApiSchoolsGetID(s))
		}
		schoolsUpdate := schools.Methods("PATCH").Subrouter()
		{
			schoolsUpdate.HandleFunc("", handles.ApiSchoolsPatch(s))
		}
		schoolsDelete := schools.Methods(http.MethodDelete).Subrouter()
		{
			schoolsDelete.HandleFunc("", handles.ApiSchoolsDelete(s))
		}
	}
}
