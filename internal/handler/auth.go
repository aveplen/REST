package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aveplen/REST/internal/constants"
	"github.com/aveplen/REST/internal/models"
)

func (h *Handler) login() http.HandlerFunc {
	h.logger.Info("login route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		var authInfo models.UserAuth
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&authInfo)
		if err != nil {
			h.logger.Warnf("login handler decode json: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		userRole, err := h.service.UserService.Login(authInfo)
		if err != nil {
			h.logger.Warnf("login handler user service login: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		token, err := h.service.JWTService.GenerateToken(userRole)
		if err != nil {
			h.logger.Warnf("login handler generate token: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("successfully logged user in")
		w.Header().Add("Content-Type", "application/jwt")
		w.Write([]byte(token))
	}
}

func (h *Handler) register() http.HandlerFunc {
	h.logger.Info("register handler route initialized")
	return func(w http.ResponseWriter, r *http.Request) {
		var authInfo models.UserAuth
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&authInfo); err != nil {
			h.logger.Warn("register handler decode json: %w", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		if err := h.service.UserService.Register(authInfo); err != nil {
			if err == constants.ErrAlreadyExists {
				h.logger.Warn("register handler: user with given email already exists", err)
				w.WriteHeader(http.StatusNotAcceptable)
				// TODO: delete me
				w.Write([]byte(fmt.Sprintf("%v", err)))
				return
			}
			h.logger.Warn("register handler bad request: %w", err)
			w.WriteHeader(http.StatusBadRequest)
			// TODO: delete me
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		h.logger.Info("register handler: succes")
		w.WriteHeader(http.StatusCreated)
		// TODO: delete me
		w.Write([]byte("user registered"))
	}
}
