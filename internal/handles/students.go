package handles

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func ApiStudentsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}

func ApiStudentsGetID(logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Error("ApiStudentsGetID: 123456789")
		logger.Info("ApiStudentsGetID")
		logger.Panic("Api blah blah blah")
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

func ApiStudentsPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}
