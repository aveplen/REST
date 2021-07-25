package handles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/aveplen/REST/internal/models"
	"github.com/gorilla/mux"
)

func ApiStudentsPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Students Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Students Post <---")
		jsonDecoder := json.NewDecoder(r.Body)
		st := models.Student{}
		err := jsonDecoder.Decode(&st)
		if err != nil {
			logger.Panic(err)
			return
		}
		jsonEncoder := json.NewEncoder(w)
		err = jsonEncoder.Encode(st)
		if err != nil {
			logger.Panic(err)
			return
		}
		// io.WriteString(w, "Hello!")
	}
}

func ApiStudentsGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Students Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Students Get <---")
		io.WriteString(w, "Hello!")
	}
}

func ApiStudentsGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Students Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Students Get ID <---")
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			panic("ApiStudentsGetID: id not found in request")
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic("ApiStudentsGetID: id is not a valid int")
		}
		fmt.Printf("ApiStudentsGetID: %d\n", id)

		io.WriteString(w, "Hello!")
	}
}

func ApiStudentsPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Students Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Students Patch <---")
		io.WriteString(w, "Hello!")
	}
}

func ApiStudentsDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Students Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Students Delete <---")
		io.WriteString(w, "Hello!")
	}
}
