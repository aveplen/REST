package handles

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ApiScoresDeleteID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Scores Delete ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Scores Delete ID <---")
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Fatal("ApiScoresDeleteID: id not found in request")
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic("ApiScoresGetID: id is not a valid int")
		}
		fmt.Printf("ApiScoresGetID: %d\n", id)
		io.WriteString(w, "Hello!")
	}
}
