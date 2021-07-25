package handles

import (
	"net/http"
)

func ApiCridentialsPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cridentials Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Cridentiala Post <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cridentials Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Cridentiala Get <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cridentials Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Cridentiala Get ID <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cridentials Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Cridentiala Patch <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiCridentialsDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cridentials Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Cridentiala Delete <---")
		// io.WriteString(w, "Hello!")
	}
}
