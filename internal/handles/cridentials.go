package handles

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ApiCridentialsDeleteID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cridentials Delete ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Cridentials Delete ID <---")
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Fatal("ApiCridentialsDeleteID: id not found in request")
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic("ApiCridentialsGetID: id is not a valid int")
		}
		fmt.Printf("ApiCridentialsGetID: %d\n", id)
		io.WriteString(w, "Hello!")
	}
}
