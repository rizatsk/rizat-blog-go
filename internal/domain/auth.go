package domain

import "context"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type Login struct {
	TokenString        string
	RefreshTokenString string
}

type AuthService interface {
	Login(ctx context.Context, req *LoginRequest) (*Login, error)
}
