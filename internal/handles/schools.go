package handles

import (
	"net/http"
)

func ApiSchoolsPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello!")
	}
}
