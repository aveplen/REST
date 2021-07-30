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

// 	Этот кусок кода закомментирован, потому что нет
// 	смысле добавлять личную информацию, которая не
// 	привязана ни к какому пользователю.

/*
func ApiCridentialsPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Cridentials()
	logger.Info("Api Cridentials Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var cridentials models.CridentialsInsert
		decoder.Decode(&cridentials)
		if err := repository.Insert(&cridentials); err != nil {
			logger.Warnf("cridentials insert: %v", err)
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
		logger.Info("cridentials insert 200 - OK")
	}
}
*/

func ApiCridentialsGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Cridentials()
	logger.Info("Api Cridentials Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		cridentials, err := repository.GetAll()
		if err != nil {
			logger.Warnf("cridentials get all: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something bad happened"))
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(cridentials)
		w.WriteHeader(http.StatusOK)
		logger.Info("cridentials get all 200 - OK")
	}
}

func ApiCridentialsGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Cridentials()
	logger.Info("Api Cridentials Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("cridentials get id: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/cridentials/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("cridentials get id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		cridentials, err := repository.GetID(id)
		if err != nil {
			logger.Warnf("cridentials get id: %v", err)
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
		encoder := json.NewEncoder(w)
		encoder.Encode(cridentials)
		w.WriteHeader(http.StatusOK)
		logger.Infof("cridentials get id(%d) all 200 - OK", id)
	}
}

func ApiCridentialsPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Cridentials()
	logger.Info("Api Cridentials Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var cridentials models.CridentialsUpdate
		decoder.Decode(&cridentials)
		if err := repository.Update(&cridentials); err != nil {
			logger.Warnf("cridentials update: %v", err)
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
		logger.Info("cridentials update 200 - OK")
	}
}

func ApiCridentialsDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Cridentials()
	logger.Info("Api Cridentials Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("cridentials delete: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/cridentials/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("cridentials delete: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		if err := repository.Delete(id); err != nil {
			logger.Warnf("cridentials delete id: %v", err)
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
		logger.Infof("cridentials delete id(%d) all 200 - OK", id)
	}
}
