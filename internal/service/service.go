package service

import (
	"github.com/aveplen/REST/internal/config"
	"github.com/aveplen/REST/internal/store"
)

type Service struct {
	UserService    *UserService
	StudentService *StudentService
	JWTService     *JWTService
}

func NewService(st *store.Store, jwtConfig config.JWT) *Service {
	jwtService := NewJWTService(jwtConfig)
	userService := NewUserService(st.Users())
	return &Service{
		UserService:    userService,
		StudentService: NewStudentService(st.Students(), userService),
		JWTService:     jwtService,
	}
}
