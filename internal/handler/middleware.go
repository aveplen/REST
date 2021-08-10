package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aveplen/REST/internal/models"
	"github.com/gorilla/mux"
)

func (h *Handler) jwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if len(header) == 0 {
				h.logger.Warn("jwt auth middleware: missing authorization header")
				w.WriteHeader(http.StatusUnauthorized)
				// TODO: delete me
				w.Write([]byte("missing authorization header"))
				return
			}
			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 {
				h.logger.Warn("jwt auth middleware: wrong amount of header parts")
				w.WriteHeader(http.StatusUnauthorized)
				// TODO: delete me
				w.Write([]byte("wrong amount of header parts"))
				return
			}
			if headerParts[0] != "Bearer" {
				h.logger.Warn("jwt auth middleware: wrong authorization scheme")
				w.WriteHeader(http.StatusUnauthorized)
				// TODO: delete me
				w.Write([]byte("wrong authorization scheme"))
				return
			}
			token := headerParts[1]
			userRole, err := h.service.JWTService.ParseToken(token)
			if err != nil {
				h.logger.Warnf("jwt auth middleware: parse token err: %v", err)
				w.WriteHeader(http.StatusUnauthorized)
				// TODO: delete me
				w.Write([]byte(fmt.Sprintf("%v", err)))
				return
			}
			h.logger.Info("jwt auth middleware: success")
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user_role", userRole)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func (h *Handler) adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			userRole, ok := ctx.Value("user_role").(*models.UserRole)
			if !ok {
				h.logger.Warn("admin only: type assertion failure")
				w.WriteHeader(http.StatusBadRequest)
				// TODO: delete me
				w.Write([]byte("type assertion failure"))
				return
			}
			if userRole.Role != "admin" {
				h.logger.Warn("admin only: user is not admin")
				w.WriteHeader(http.StatusForbidden)
				// TODO: delete me
				w.Write([]byte("not enough rights"))
				return
			}
			h.logger.Info("admin only middleware: success")
			next.ServeHTTP(w, r)
		})
}

func (h *Handler) adminOrUser(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			muxVars := mux.Vars(r)
			urlUserID, ok := muxVars["user_id"]
			if !ok {
				h.logger.Warn("admin or user middleware: no user_id in URL")
				w.WriteHeader(http.StatusBadRequest)
				// TODO: delete me
				w.Write([]byte("no user id in URL"))
				return
			}
			intUrlUserID, err := strconv.ParseInt(urlUserID, 10, 64)
			if err != nil {
				h.logger.Warnf("admin or user middleware: parse int err: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				// TODO: delete me
				w.Write([]byte(fmt.Sprintf("%v", err)))
				return
			}
			ctx := r.Context()
			jwtUserRole, ok := ctx.Value("user_role").(*models.UserRole)
			if !ok {
				h.logger.Warn("admin or user middleware: type assertion failure")
				w.WriteHeader(http.StatusInternalServerError)
				// TODO: delete me
				w.Write([]byte("type assertion failure"))
				return
			}
			if !(intUrlUserID == jwtUserRole.UserID || jwtUserRole.Role == "admin") {
				h.logger.Warn("admin or user middleware: not user or admin")
				w.WriteHeader(http.StatusForbidden)
				// TODO: delete me
				w.Write([]byte("not admin or user"))
				return
			}
			h.logger.Info("admin or user middleware: success")
			next.ServeHTTP(w, r)
		})
}
