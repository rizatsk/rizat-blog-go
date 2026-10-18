package domain

import (
	"context"
)

type AuthToken struct {
	ID        string `json:"id"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type AuthTokenRepository interface {
	GetAuthTokenByUserId(ctx context.Context, userId string) ([]AuthToken, error)
	InsertToken(ctx context.Context, id string, userId string) error
	DeleteTokenByIds(ctx context.Context, ids []string) error
}

func (AuthToken) TableName() string {
	return "rizat_blog.auth_tokens"
}
