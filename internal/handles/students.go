package handles

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aveplen/REST/internal/models"
	"github.com/gorilla/mux"
)

func ApiStudentsPost(s IServer) http.HandlerFunc {
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
		// decode request body
		decoder := json.NewDecoder(r.Body)
		var student models.StudentInsert
		decoder.Decode(&student)

		// try to insert value decoded from request body into table
		// check if something bad happened
		if err := repository.Insert(&student); err != nil {
			logger.Warnf("students insert handle: %v", err)
			switch errors.Unwrap(err) {
			// if db returned error that correspondes to zero
			// rows inserted then it is a bad request
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("check request json"))
				logger.Warnf("%d, %s", http.StatusBadRequest, "check request json")
			// all other cases are internal server error
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something bad happened"))
				logger.Warnf("%d, %s", http.StatusInternalServerError, "something bad happened")
			}
			// stop function execution to avoid else keyword
			return
		}
		// response 200 and log success
		w.WriteHeader(http.StatusOK)
		logger.Info("students insert 200 - OK")
	}
}

func ApiStudentsGet(s IServer) http.HandlerFunc {
	// check that logger is not nil
	logger := s.GetLogger()
	if logger == nil {
		panic(fmt.Errorf("students get: logger is nil, can't proceed"))
	}
	// check that store is not nil
	store := s.GetStore()
	if store == nil {
		logger.Fatal(fmt.Errorf("students get: store is nil, can't proceed"))
	}
	// check that repository is not nil
	repository := store.Student()
	if repository == nil {
		logger.Fatal(fmt.Errorf("students get: repository is nil, can't proceed"))
	}
	// log success
	logger.Info("Api Students Get route initialized")

	// return handler as an anonymous function
	return func(w http.ResponseWriter, r *http.Request) {
		// get all students from the repository
		//
		// if db returns error then always internal server error
		// because request is just a plain blank GET with no body
		students, err := repository.GetAll()
		if err != nil {
			logger.Warnf("students get all: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something bad happened"))
			// stop function execution to avoid else keyword
			return
		}
		// encode the result into json and attach in to response
		encoder := json.NewEncoder(w)
		encoder.Encode(students)
		// response 200 and log success
		w.WriteHeader(http.StatusOK)
		logger.Info("students get all 200 - OK")
	}
}

func ApiStudentsGetID(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Student()
	logger.Info("Api Students Get ID route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("students get id: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/schools/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("students get id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		student, err := repository.GetID(id)
		if err != nil {
			logger.Warnf("students get id: %v", err)
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
		encoder.Encode(student)
		w.WriteHeader(http.StatusOK)
		logger.Infof("students get id(%d) all 200 - OK", id)
	}
}

func ApiStudentsPatch(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Student()
	logger.Info("Api Students Patch route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var student models.StudentUpdate
		decoder.Decode(&student)
		if err := repository.Update(&student); err != nil {
			logger.Warnf("students update: %v", err)
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
		logger.Info("students update 200 - OK")
	}
}

func ApiStudentsDelete(s IServer) http.HandlerFunc {
	logger := s.GetLogger()
	repository := s.GetStore().Student()
	logger.Info("Api Students Delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			logger.Warn("students delete: id not found in request")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id not found in request\nshould be: /api/students/{id:[0-9]+}"))
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			logger.Warnf("students delete: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("requested id is not a valid int (probably too large)"))
			return
		}
		if err := repository.Delete(id); err != nil {
			logger.Warnf("students delete id: %v", err)
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
		logger.Infof("students delete id(%d) 200 - OK", id)
	}
}
