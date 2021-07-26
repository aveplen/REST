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

func ApiCitiesPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Post route initialized")
	repository := s.GetStore().City()
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Post <---")
		decoder := json.NewDecoder(r.Body)
		city := &models.City{}
		decoder.Decode(&city)
		city, err := repository.Insert(city)
		if err != nil {
			logger.Fatal(err)
		}
		io.WriteString(w, fmt.Sprintf("%d", city.CityID))
	}
}

func ApiCitiesGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Get route initialized")
	repository := s.GetStore().City()
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Get <---")
		cities, err := repository.GetAll()
		if err != nil {
			logger.Panicf("error from db:\n%s\n", err)
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(cities)
	}
}

func ApiCitiesGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Get ID route initialized")
	repository := s.GetStore().City()
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("---> Api Cities Get ID <---")
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Fatal("ApiCitiesDeleteID: id not found in request")
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Fatal("ApiCitiesGetID: id is not a valid int")
		}
		city, err := repository.GetID(id)
		if err != nil {
			logger.Fatal(err)
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(city)
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

func ApiCitiesDeleteID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	logger.Info("Api Cities Delete ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("---> Api Cities Delete ID <---")
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Fatal("ApiCitiesDeleteID: id not found in request")
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic("ApiCitiesGetID: id is not a valid int")
		}
		fmt.Printf("ApiCitiesGetID: %d\n", id)
		io.WriteString(w, "Hello!")
	}
}
