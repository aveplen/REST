package handles

import (
	"net/http"
)

func ApiScoresPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}
