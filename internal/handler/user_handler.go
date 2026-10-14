package handler

import (
	"net/http"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"
	"service_blog_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserService
	env         *config.Environment
}

func NewUserHandler(userService domain.UserService, env *config.Environment) *UserHandler {
	return &UserHandler{
		userService: userService,
		env:         env,
	}
}

func (user *UserHandler) GetUserByToken(ctx *gin.Context) {
	dataUser := user.userService.GetUser(ctx)
	response.Success(
		ctx,
		http.StatusOK,
		"Success get data user",
		dataUser,
	)
}
