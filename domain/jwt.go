package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	Handle string `json:"handle"`
	ID     string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
