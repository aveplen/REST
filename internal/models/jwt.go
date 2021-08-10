package models

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	jwt.StandardClaims
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
}
