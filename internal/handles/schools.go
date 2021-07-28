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

func ApiSchoolsPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().School()
	logger.Info("Api Schools Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var school models.SchoolInsert
		decoder.Decode(&school)
		if err := repository.Insert(&school); err != nil {
			logger.Warnf("schools insert: %v", err)
			switch errors.Unwrap(err) {
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("check request json"))
				logger.Warnf("%d, %s", http.StatusBadRequest, "check request json")
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something bad happened"))
				logger.Warnf("%d, %s", http.StatusInternalServerError, "something bad happened")
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		logger.Info("schools insert 200 - OK")
	}
}

func ApiSchoolsGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().School()
	logger.Info("Api Schools Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		schools, err := repository.GetAll()
		if err != nil {
			logger.Warnf("schools get all: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something bad happened"))
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(schools)
		w.WriteHeader(http.StatusOK)
		logger.Info("schools get all 200 - OK")
	}
}

func ApiSchoolsGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().School()
	logger.Info("Api Schools Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("schools get id: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/schools/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("schools get id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		school, err := repository.GetID(id)
		if err != nil {
			logger.Warnf("schools get id: %v", err)
			switch errors.Unwrap(err) {
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("no entry with such index"))
				logger.Warnf("%d, %s", http.StatusBadRequest, "check request json")
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something bad happened"))
				logger.Warnf("%d, %s", http.StatusInternalServerError, "something bad happened")
			}
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(school)
		w.WriteHeader(http.StatusOK)
		logger.Infof("schools get id(%d) all 200 - OK", id)
	}
}

func ApiSchoolsPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().School()
	logger.Info("Api Schools Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var school models.SchoolUpdate
		decoder.Decode(&school)
		if err := repository.Update(&school); err != nil {
			logger.Warnf("schools update: %v", err)
			switch errors.Unwrap(err) {
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("check request json"))
				logger.Warnf("%d, %s", http.StatusBadRequest, "check request json")
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something bad happened"))
				logger.Warnf("%d, %s", http.StatusInternalServerError, "something bad happened")
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		logger.Info("schools update 200 - OK")
	}
}

func ApiSchoolsDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().School()
	logger.Info("Api Schools Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("schools delete: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/schools/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("schools delete: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		if err := repository.Delete(id); err != nil {
			logger.Warnf("schools delete id: %v", err)
			switch errors.Unwrap(err) {
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("no entry with such index"))
				logger.Warnf("%d, %s", http.StatusBadRequest, "no entry with such index")
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something bad happened"))
				logger.Warnf("%d, %s", http.StatusInternalServerError, "something bad happened")
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		logger.Infof("schools delete id(%d) 200 - OK", id)
	}
}
