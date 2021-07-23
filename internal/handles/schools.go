package handles

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func ApiSchoolsPost(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Schools Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsGet(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Schools Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsGetID(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Schools Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsPatch(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Schools Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsDelete(logger *logrus.Logger) http.HandlerFunc {
	logger.Info("Api Schools Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}
