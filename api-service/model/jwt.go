package model

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserID int64  `json:"user_id"`
	Roles  string `json:"roles"`
	jwt.RegisteredClaims
}
