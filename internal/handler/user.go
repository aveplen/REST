package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aveplen/REST/internal/models"
	"github.com/gorilla/mux"
)

func (h *Handler) userGet() http.HandlerFunc {
	h.logger.Info("user get route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, ok := ctx.Value("user_role").(*models.UserRole)
		if !ok {
			h.logger.Warn("user get handler: no user_role value in context")
			w.WriteHeader(http.StatusUnauthorized)
			// TODO: delete me
			w.Write([]byte("type assertion failure"))
			return
		}
		muxVars := mux.Vars(r)
		userID, ok := muxVars["user_id"]
		if !ok {
			h.logger.Warn("user get handler: no user_id in URL")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("no user id in URL"))
			return
		}
		integerUserID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			h.logger.Warn("user get handler: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		userResponse, err := h.service.UserService.GetProfile(integerUserID)
		if err != nil {
			h.logger.Warnf("user get handler: user service get profile: %v", err)
			w.WriteHeader(http.StatusNotFound)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("user get handler: success")
		w.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(userResponse)
	}
}

func (h *Handler) userDelete() http.HandlerFunc {
	h.logger.Info("user delete route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		muxVars := mux.Vars(r)
		userID, ok := muxVars["user_id"]
		if !ok {
			h.logger.Warn("user delete handler: no user_id in URL")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("no user id in URL"))
			return
		}
		integerUserID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			h.logger.Warn("user delete handler: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		if err := h.service.UserService.Delete(integerUserID); err != nil {
			h.logger.Warnf("user delete handler: user service delete: %v", err)
			w.WriteHeader(http.StatusNotFound)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("user delete handler: success")
		// TODO: delete me
		w.Write([]byte("user deleted"))
	}
}

func (h *Handler) userUpdate() http.HandlerFunc {
	h.logger.Info("user update route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		muxVars := mux.Vars(r)
		userID, ok := muxVars["user_id"]
		if !ok {
			h.logger.Warn("user update handler: no user_id in URL")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte("no useer is in URL"))
			return
		}
		integerUserID, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			h.logger.Warn("user update handler: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		decoder := json.NewDecoder(r.Body)
		var user models.UserUpdate
		if err := decoder.Decode(&user); err != nil {
			h.logger.Warn("user update handler: request body is not valid")
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		if integerUserID != user.UserID {
			h.logger.Warn("user update handler ctx.user_id != URL.user_id")
			w.WriteHeader(http.StatusForbidden)
			// TODO: delete me
			w.Write([]byte("context user_id != URL user_id"))
			return
		}
		if err := h.service.UserService.Update(user); err != nil {
			h.logger.Warnf("user update handler: user service update: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("user update handler: success")
		// TODO: delete me
		w.Write([]byte("user updated"))
	}
}
