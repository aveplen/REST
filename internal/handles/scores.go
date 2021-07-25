package handles

import (
	"net/http"
)

func ApiScoresPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Scores Post <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Scores Get <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Scores Get ID <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Scores Patch <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Scores Delete <---")
		// io.WriteString(w, "Hello!")
	}
}
