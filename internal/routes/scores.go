package routes

import (
	"net/http"

	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteScores(parent *mux.Router, s handles.IServer) {
	scores := parent.PathPrefix("/scores").Subrouter()
	{
		scoresCreate := scores.Methods("POST").Subrouter()
		{
			scoresCreate.HandleFunc("", handles.ApiScoresPost(s))
		}
		scoresRead := scores.Methods("GET").Subrouter()
		{
			scoresRead.HandleFunc("", handles.ApiScoresGet(s))
			scoresRead.HandleFunc("/{id:[0-9]+}", handles.ApiScoresGetID(s))
		}
		scoresUpdate := scores.Methods("PATCH").Subrouter()
		{
			scoresUpdate.HandleFunc("", handles.ApiScoresPatch(s))
		}
		scoresDelete := scores.Methods(http.MethodDelete).Subrouter()
		{
			scoresDelete.HandleFunc("", handles.ApiScoresDelete(s))
		}
	}
}
