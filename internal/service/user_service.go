package service

import (
	"context"
	"errors"
	"log"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"
	"service_blog_go/pkg/asynclock"
	"service_blog_go/pkg/tokenize"
	"uuid"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo domain.UserRepository
	authRepo domain.AuthTokenRepository
	env      *config.Environment
	syncLock *asynclock.SyncLock
}

func NewUserService(userRepo domain.UserRepository, authRepo domain.AuthTokenRepository, env *config.Environment, syncLock *asynclock.SyncLock) domain.UserService {
	return &UserService{
		userRepo: userRepo,
		authRepo: authRepo,
		env:      env,
		syncLock: syncLock,
	}
}

func (userService *UserService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.Login, error) {
	user, err := userService.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("User is not registered, error: %v", err)
		return nil, errors.New("Email or password is wrong")
	}

	err_validate_pass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err_validate_pass != nil {
		log.Printf("Password is wrong, err : %v", err_validate_pass)
		return nil, errors.New("Email or password is wrong")
	}

	tokenString, err := tokenize.GenerateAccessToken(userService.env, user)
	if err != nil {
		log.Printf("Failed generate token jwt error: %v", err)
		return nil, errors.New("Failed generate token jwt")
	}

	authTokenId := uuid.NewV7().String()
	refreshTokenString, err := tokenize.GenerateRefreshToken(userService.env, authTokenId, user.ID)
	if err != nil {
		log.Printf("Failed generate refersh token jwt error: %v", err)
		return nil, errors.New("Failed generate refresh token jwt")
	}

	// Locking
	userService.syncLock.Lock(req.Email)
	defer userService.syncLock.Unlock(req.Email)

	authTokens, err := userService.authRepo.GetAuthTokenByUserId(ctx, user.ID)
	if err != nil {
		log.Printf("Failed get tokens user from table : %v", err)
		return nil, errors.New("Failed get tokens user from table")
	}

	if len(authTokens) >= userService.env.UserMaxDevice {
		var idsTokenOld []string
		for i := userService.env.UserMaxDevice - 1; i < len(authTokens); i++ {
			idsTokenOld = append(idsTokenOld, authTokens[i].ID)
		}

		var errDeleteToken error
		errDeleteToken = userService.authRepo.DeleteTokenByIds(ctx, idsTokenOld)
		if errDeleteToken != nil {
			log.Printf("Failed delete token by id : %v", errDeleteToken)
			return nil, errors.New("Failed delete token by id")
		}
	}

	var errInsertToken error
	errInsertToken = userService.authRepo.InsertToken(ctx, authTokenId, user.ID)
	if errInsertToken != nil {
		log.Printf("Failed create token from table : %v", errInsertToken)
		return nil, errors.New("Failed create token from table")
	}

	return &domain.Login{TokenString: tokenString, RefreshTokenString: refreshTokenString}, nil
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
