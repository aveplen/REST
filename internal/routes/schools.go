package routes

import (
	"net/http"

	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func RouteSchools(parent *mux.Router, logger *logrus.Logger) {
	schools := parent.PathPrefix("/schools").Subrouter()
	{
		schoolsCreate := schools.Methods("POST").Subrouter()
		{
			schoolsCreate.HandleFunc("", handles.ApiSchoolsPost(logger))
		}
		schoolsRead := schools.Methods("GET").Subrouter()
		{
			schoolsRead.HandleFunc("", handles.ApiSchoolsGet(logger))
			schoolsRead.HandleFunc("/{id:[0-9]+}", handles.ApiSchoolsGetID(logger))
		}
		schoolsUpdate := schools.Methods("PATCH").Subrouter()
		{
			schoolsUpdate.HandleFunc("", handles.ApiSchoolsPatch(logger))
		}
		schoolsDelete := schools.Methods(http.MethodDelete).Subrouter()
		{
			schoolsDelete.HandleFunc("", handles.ApiSchoolsDelete(logger))
		}
	}
}
