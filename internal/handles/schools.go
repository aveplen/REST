package handles

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ApiSchoolsPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Schools Post <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Schools Get <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Schools Get ID <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Schools Patch <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Schools Delete <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiSchoolsDeleteID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Schools Delete ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Schools Delete ID <---")
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Fatal("ApiSchoolsDeleteID: id not found in request")
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic("ApiSchoolsGetID: id is not a valid int")
		}
		fmt.Printf("ApiSchoolsGetID: %d\n", id)
		io.WriteString(w, "Hello!")
	}
}
