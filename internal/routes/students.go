package routes

import (
	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteStudents(parent *mux.Router, s handles.IServer) {
	students := parent.PathPrefix("/students").Subrouter()
	{
		studentsCreate := students.Methods("POST").Subrouter()
		{
			studentsCreate.HandleFunc("", handles.ApiStudentsPost(s))
		}
		studentsRead := students.Methods("GET").Subrouter()
		{
			studentsRead.HandleFunc("", handles.ApiStudentsGet(s))
			studentsRead.HandleFunc("/{id:[0-9]+}", handles.ApiStudentsGetID(s))
		}
		studentsUpdate := students.Methods("PATCH").Subrouter()
		{
			studentsUpdate.HandleFunc("", handles.ApiStudentsPatch(s))
		}
		studentsDelete := students.Methods("DELETE").Subrouter()
		{
			studentsDelete.HandleFunc("", handles.ApiStudentsDelete(s))
			studentsDelete.HandleFunc("/{id:[0-9]+}", handles.ApiStudentsDeleteID(s))
		}
	}
}
