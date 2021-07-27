package handles

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/aveplen/REST/internal/models"
	"github.com/gorilla/mux"
)

func ApiCitiesPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().City()
	logger.Info("Api Cities Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var city models.CityInsert
		decoder.Decode(&city)
		if err := repository.Insert(&city); err != nil {
			logger.Warnf("cities insert: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something bad happened"))
			return
		}
		w.WriteHeader(http.StatusOK)
		logger.Info("cities insert 200 - OK")
	}
}

func ApiCitiesGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().City()
	logger.Info("Api Cities Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		cities, err := repository.GetAll()
		if err != nil {
			logger.Warnf("cities get all: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something bad happened"))
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(cities)
		w.WriteHeader(http.StatusOK)
		logger.Info("cities get all 200 - OK")
	}
}

func ApiCitiesGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().City()
	logger.Info("Api Cities Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("cities get id: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/cities/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("cities get id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		city, err := repository.GetID(id)
		if err != nil {
			logger.Warnf("cities get id: %v", err)
			err := errors.Unwrap(err)
			switch err {
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("no entry with such index"))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something bad happened"))
			}
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(city)
		w.WriteHeader(http.StatusOK)
		logger.Infof("cities get id(%d) all 200 - OK", id)
	}
}

func ApiCitiesPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().City()
	logger.Info("Api Cities Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var city models.CityUpdate
		decoder.Decode(&city)
		if err := repository.Update(&city); err != nil {
			logger.Warnf("cities update: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something bad happened"))
			return
		}
		w.WriteHeader(http.StatusOK)
		logger.Info("cities update 200 - OK")
	}
}

func ApiCitiesDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().City()
	logger.Info("Api Cities Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("cities delete: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/cities/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("cities delete: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		if repository.Delete(id); err != nil {
			logger.Warnf("cities delete id: %v", err)
			switch err {
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("no entry with such index"))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something bad happened"))
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		logger.Infof("cities delete id(%d) all 200 - OK", id)
	}
}
