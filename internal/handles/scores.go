package handles

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func ApiScoresPost(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Scores Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresGet(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Scores Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresGetID(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Scores Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresPatch(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Scores Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiScoresDelete(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Scores Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}
