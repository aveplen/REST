package routes

import (
	"net/http"

	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func RouteScores(parent *mux.Router, logger *logrus.Logger) {
	scores := parent.PathPrefix("/scores").Subrouter()
	{
		scoresCreate := scores.Methods("POST").Subrouter()
		{
			scoresCreate.HandleFunc("", handles.ApiScoresPost(logger))
		}
		scoresRead := scores.Methods("GET").Subrouter()
		{
			scoresRead.HandleFunc("", handles.ApiScoresGet(logger))
			scoresRead.HandleFunc("/{id:[0-9]+}", handles.ApiScoresGetID(logger))
		}
		scoresUpdate := scores.Methods("PATCH").Subrouter()
		{
			scoresUpdate.HandleFunc("", handles.ApiScoresPatch(logger))
		}
		scoresDelete := scores.Methods(http.MethodDelete).Subrouter()
		{
			scoresDelete.HandleFunc("", handles.ApiScoresDelete(logger))
		}
	}
}
