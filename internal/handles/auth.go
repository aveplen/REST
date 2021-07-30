package handles

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aveplen/REST/internal/models"
)

func ApiAuthLogin(s IServer) http.HandlerFunc {
	// check that logger is not nil
	logger := s.GetLogger()
	if logger == nil {
		panic(fmt.Errorf("logger is nil, can't proceed"))
	}
	// check that store is not nil
	store := s.GetStore()
	if store == nil {
		logger.Panic(fmt.Errorf("store is nil, can't proceed"))
	}
	// check that repository is not nil
	repository := store.Student()
	if repository == nil {
		logger.Panic(fmt.Errorf("repository is nil, can't proceed"))
	}
	// log success
	logger.Info("Api Students Post route initialized")

	// return handler as an anonymous function
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var city models.CityInsert
		decoder.Decode(&city)
		if err := repository.Insert(&city); err != nil {
			logger.Warnf("cities insert: %v", err)
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
		logger.Info("cities insert 200 - OK")
	}
}

func ApiAuthRegister(s IServer) http.HandlerFunc {
	// check that logger is not nil
	logger := s.GetLogger()
	if logger == nil {
		panic(fmt.Errorf("logger is nil, can't proceed"))
	}
	// check that store is not nil
	store := s.GetStore()
	if store == nil {
		logger.Panic(fmt.Errorf("store is nil, can't proceed"))
	}
	// check that repository is not nil
	repository := store.Student()
	if repository == nil {
		logger.Panic(fmt.Errorf("repository is nil, can't proceed"))
	}
	// log success
	logger.Info("Api Students Post route initialized")

	// return handler as an anonymous function
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var city models.CityInsert
		decoder.Decode(&city)
		if err := repository.Insert(&city); err != nil {
			logger.Warnf("cities insert: %v", err)
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
		logger.Info("cities insert 200 - OK")
	}
}
