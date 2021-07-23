package handles

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func ApiCitiesPost(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cities Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCitiesGet(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cities Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCitiesGetID(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cities Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCitiesPatch(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cities Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCitiesDelete(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cities Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}
