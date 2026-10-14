package handler

import (
	"net/http"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"
	"time"

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
	ctx.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"message":    "Success get dat user",
		"data":       dataUser,
		"created_at": time.Now().Format(user.env.FormatDateTime),
	})
}
