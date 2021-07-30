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

// 	Я не знаю, зачем вообще написал этот кусок.
// 	Юзкейсов у него нету (нет смысла добавлять
//	набор результатов, который никому не принадлежит),
//	но удалять его как-то жалко, поэтому пусть остаётся
// 	пока.

/*
func ApiScoresPost(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Score()
	logger.Info("Api Scores Post route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var score models.ScoreInsert
		decoder.Decode(&score)
		if err := repository.Insert(&score); err != nil {
			logger.Warnf("scores insert: %v", err)
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
		logger.Info("scores insert 200 - OK")
	}
}
*/

func ApiScoresGet(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Score()
	logger.Info("Api Scores Get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		scores, err := repository.GetAll()
		if err != nil {
			logger.Warnf("scores get all: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something bad happened"))
			return
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(scores)
		w.WriteHeader(http.StatusOK)
		logger.Info("scores get all 200 - OK")
	}
}

func ApiScoresGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Score()
	logger.Info("Api Scores Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("scores get id: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/scores/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("scores get id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		score, err := repository.GetID(id)
		if err != nil {
			logger.Warnf("scores get id: %v", err)
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
		encoder.Encode(score)
		w.WriteHeader(http.StatusOK)
		logger.Infof("scores get id(%d) all 200 - OK", id)
	}
}

func ApiScoresPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Score()
	logger.Info("Api Scores Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var score models.ScoreUpdate
		decoder.Decode(&score)
		if err := repository.Update(&score); err != nil {
			logger.Warnf("scores update: %v", err)
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
		logger.Info("scores update 200 - OK")
	}
}

func ApiScoresDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Score()
	logger.Info("Api Scores Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("scores delete: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/scores/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("scores delete: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		if err := repository.Delete(id); err != nil {
			logger.Warnf("scores delete id: %v", err)
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
		logger.Infof("scores delete id(%d) all 200 - OK", id)
	}
}
