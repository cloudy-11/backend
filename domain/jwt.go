package domain

import "github.com/golang-jwt/jwt/v4"

const (
	USER_ROLE  = "user"
	ADMIN_ROLE = "admin"
	GUEST_ROLE = "guest"
)

type JwtCustomClaims struct {
	Handle string `json:"handle"`
	ID     string `json:"id"`
	Role   string `json:"role"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
