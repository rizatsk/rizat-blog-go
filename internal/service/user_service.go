package service

import (
	"context"
	"errors"
	"log"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo domain.UserRepository
	env      *config.Environment
}

func NewUserService(userRepo domain.UserRepository, env *config.Environment) domain.UserService {
	return &UserService{
		userRepo: userRepo,
		env:      env,
	}
}

type JWTClaims struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	jwt.RegisteredClaims
}

func (userService *UserService) Login(ctx context.Context, req *domain.LoginRequest) (string, error) {
	user, err := userService.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("User is not registered, error: %v", err)
		return "", errors.New("Email or password is wrong")
	}

	err_validate_pass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err_validate_pass != nil {
		log.Printf("Password is wrong, err : %v", err_validate_pass)
		return "", errors.New("Email or password is wrong")
	}

	claims := JWTClaims{
		UserID:    user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(userService.env.FormatDateTime),
		UpdatedAt: user.UpdatedAt.Format(userService.env.FormatDateTime),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(userService.env.JWTExpiredHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(userService.env.JWTSecretKey))
	if err != nil {
		log.Printf("Failed generate token jwt error: %v", err)
		return "", errors.New("Failed generate token jwt")
	}

	return tokenString, nil
}

func (userService UserService) GetUser(ctx *gin.Context) *domain.GetUser {
	return &domain.GetUser{
		UserId:    ctx.MustGet("user_id").(string),
		Name:      ctx.MustGet("name").(string),
		Email:     ctx.MustGet("email").(string),
		Role:      ctx.MustGet("role").(string),
		CreatedAt: ctx.MustGet("created_at").(string),
		UpdatedAt: ctx.MustGet("updated_at").(string),
	}
}
