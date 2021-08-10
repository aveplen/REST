package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aveplen/REST/internal/models"
)

func (h *Handler) studentAttach() http.HandlerFunc {
	h.logger.Info("student attach route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userRole, ok := ctx.Value("user_role").(*models.UserRole)
		if !ok {
			h.logger.Warn("student attach handler: type assertion failure")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("type assertion failure"))
			return
		}
		var student models.StudentInsert
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&student); err != nil {
			h.logger.Warnf("student attach handler: json body decode err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		if err := h.service.StudentService.AttachStudent(userRole.UserID, student); err != nil {
			h.logger.Warnf("student attach handler: attach student err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("student attach handler: success")
		w.WriteHeader(http.StatusCreated)
		// TODO: delete me
		w.Write([]byte("student attached"))
	}
}

func (h *Handler) studentDetach() http.HandlerFunc {
	h.logger.Info("student detach route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userRole, ok := ctx.Value("user_role").(*models.UserRole)
		if !ok {
			h.logger.Warn("student detach handler: type assertion failure")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("type assertion failure"))
			return
		}
		if err := h.service.StudentService.DetachStudent(userRole.UserID); err != nil {
			h.logger.Warnf("student detach handler: detach student err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("student detach handler: success")
		// TODO: delete me
		w.Write([]byte("student detached"))
	}
}

func (h *Handler) studentGet() http.HandlerFunc {
	h.logger.Info("student get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userRole, ok := ctx.Value("user_role").(*models.UserRole)
		if !ok {
			h.logger.Warn("student get handler: type assertion failure")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("type assertion failure"))
			return
		}
		student, err := h.service.StudentService.GetStudentInfoFromUserID(userRole.UserID)
		if err != nil {
			h.logger.Warnf("student get handler: get student info from user id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("student get handler: success")
		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(student)
	}
}

func (h *Handler) studentUpdate() http.HandlerFunc {
	h.logger.Info("student update route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, ok := ctx.Value("user_role").(*models.UserRole)
		if !ok {
			h.logger.Warn("student update handler: type assertion failure")
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: delete me
			w.Write([]byte("type assertion failure"))
			return
		}
		var student models.StudentUpdate
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&student); err != nil {
			h.logger.Warnf("student update handler: json body decode: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		if err := h.service.StudentService.UpdateStudent(student); err != nil {
			h.logger.Warnf("student update handler: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("student update handler: success")
		// TODO: delete me
		w.Write([]byte("student updated"))
	}
}
