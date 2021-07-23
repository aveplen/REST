package handles

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func ApiCridentialsPost(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cridentials Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsGet(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cridentials Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsGetID(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cridentials Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsPatch(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cridentials Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsDelete(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Cridentials Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}
