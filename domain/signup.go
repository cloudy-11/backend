package domain

import (
	"context"
)

type SignupRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUseCase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
