package routes

import (
	"net/http"

	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func RouteStudents(parent *mux.Router, logger *logrus.Logger) {
	students := parent.PathPrefix("/students").Subrouter()
	{
		studentsCreate := students.Methods("POST").Subrouter()
		{
			studentsCreate.HandleFunc("", handles.ApiStudentsPost(logger))
		}
		studentsRead := students.Methods("GET").Subrouter()
		{
			studentsRead.HandleFunc("", handles.ApiStudentsGet(logger))
			studentsRead.HandleFunc("/{id:[0-9]+}", handles.ApiStudentsGetID(logger))
		}
		studentsUpdate := students.Methods("PATCH").Subrouter()
		{
			studentsUpdate.HandleFunc("", handles.ApiStudentsPatch(logger))
		}
		studentsDelete := students.Methods(http.MethodDelete).Subrouter()
		{
			studentsDelete.HandleFunc("", handles.ApiStudentsDelete(logger))
		}
	}
}
