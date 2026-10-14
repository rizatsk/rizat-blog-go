package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `gorm:"primaryKey;" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"` // Hidden dari JSON response
	Role      string    `gorm:"type:varchar(20);default:'user'" json:"role"`
	Status    string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type GetUser struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type UserService interface {
	Login(ctx context.Context, req *LoginRequest) (string, error)
	GetUser(ctx *gin.Context) *GetUser
}

func (User) TableName() string {
	return "rizat_blog.users"
}

func (user *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}
