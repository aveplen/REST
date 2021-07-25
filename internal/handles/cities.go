package handles

import (
	"encoding/json"
	"net/http"
)

func ApiCitiesPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Post <---")
	}
}

func ApiCitiesGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Get route initialized")
	repository := s.GetStore().City()
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Get <---")
		res, err := repository.GetAll()
		if err != nil {
			logger.Panicf("error from db:\n%s\n", err)
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(res)
	}
}

func ApiCitiesGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Get ID <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiCitiesPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Patch <---")
		// io.WriteString(w, "Hello!")
	}
}

func ApiCitiesDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Delete <---")
		// io.WriteString(w, "Hello!")
	}
}
