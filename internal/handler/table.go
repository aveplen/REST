package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aveplen/REST/internal/models"
	"github.com/gorilla/mux"
)

//final [ ]
func (h *Handler) tableFull() http.HandlerFunc {
	h.logger.Info("table full route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		table, err := h.service.StudentService.GetAll()
		if err != nil {
			h.logger.Warnf("table full: get all err: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("table full handler: success")
		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string][]models.StudentResponse{
			"students": table,
		})
	}
}

func (h *Handler) tablePage() http.HandlerFunc {
	h.logger.Info("table page route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		muxVars := mux.Vars(r)
		pageSize, ok := muxVars["size"]
		if !ok {
			h.logger.Warn("table page: no page size in URL")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("no page size in URL"))
			return
		}
		intPageSize, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			h.logger.Warnf("table page: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		pageNum, ok := muxVars["page"]
		if !ok {
			h.logger.Warn("table page: no page number in URL")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("no page number in URL"))
			return
		}
		intPageNum, err := strconv.ParseInt(pageNum, 10, 64)
		if err != nil {
			h.logger.Warnf("table page: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		amntOfRecords, err := h.service.StudentService.CountAll()
		if err != nil {
			h.logger.Warnf("table page: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		maxPageNum := amntOfRecords / intPageSize
		if amntOfRecords%intPageSize != 0 {
			maxPageNum++
		}
		if intPageNum > maxPageNum {
			h.logger.Warn("table page: int page num is greater then max page num")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("page num is greater then maximum page num"))
			return
		}
		table, err := h.service.StudentService.GetPage(models.StudentPageRequest{
			PageSize: intPageSize,
			PageNum:  intPageNum,
		})
		if err != nil {
			h.logger.Warnf("table page: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("table full handler: success")
		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string][]models.StudentResponse{
			"students": table,
		})
	}
}
