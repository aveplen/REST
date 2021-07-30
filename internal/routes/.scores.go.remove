package routes

import (
	"github.com/aveplen/REST/internal/handles"
	"github.com/gorilla/mux"
)

func RouteScores(parent *mux.Router, s handles.IServer) {
	scores := parent.PathPrefix("/scores").Subrouter()
	{
		//  Этот раут закомментирован, потому что нет смысла
		//  добавлять результаты, которые никому не принадлежат,
		//  но соответствующий код обработчика и запроса в базу
		//  удалять у меня рука не поднялась, поэтому это тоже
		//  пусть остаётся.

		/*
			scoresCreate := scores.Methods("POST").Subrouter()
			{
				scoresCreate.HandleFunc("", handles.ApiScoresPost(s))
			}
		*/
		scoresRead := scores.Methods("GET").Subrouter()
		{
			scoresRead.HandleFunc("", handles.ApiScoresGet(s))
			scoresRead.HandleFunc("/{id:[0-9]+}", handles.ApiScoresGetID(s))
		}
		scoresUpdate := scores.Methods("PATCH").Subrouter()
		{
			scoresUpdate.HandleFunc("", handles.ApiScoresPatch(s))
		}
		scoresDelete := scores.Methods("DELETE").Subrouter()
		{
			scoresDelete.HandleFunc("/{id:[0-9]+}", handles.ApiScoresDelete(s))
		}
	}
}
