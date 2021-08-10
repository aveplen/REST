package service

import (
	"fmt"
	"time"

	"github.com/aveplen/REST/internal/config"
	"github.com/aveplen/REST/internal/models"
	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	cfg config.JWT
}

func NewJWTService(cfg config.JWT) *JWTService {
	return &JWTService{
		cfg: cfg,
	}
}

func (j *JWTService) GenerateToken(user models.UserRole) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.cfg.Expire).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: user.UserID,
		Role:   user.Role,
	})
	res, err := token.SignedString([]byte(j.cfg.Key))
	if err != nil {
		return "", fmt.Errorf("jwt service generatetoken: %w", err)
	}
	return res, nil
}

func (j *JWTService) ParseToken(tok string) (*models.UserRole, error) {
	token, err := jwt.ParseWithClaims(tok, &models.JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", fmt.Errorf("doesn't ever trigger lol")
			}
			return []byte(j.cfg.Key), nil
		})
	if err != nil {
		return nil, fmt.Errorf("jwt service parse token: %w", err)
	}
	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return &models.UserRole{
			UserID: claims.UserID,
			Role:   claims.Role,
		}, nil
	}
	return nil, fmt.Errorf("jwt service parse token: type assertion failure")

}
