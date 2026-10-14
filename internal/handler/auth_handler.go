package handler

import (
	"net/http"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService domain.UserService
	env         *config.Environment
}

func NewAuthHandler(userService domain.UserService, env *config.Environment) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		env:         env,
	}
}

func (authHandler *AuthHandler) Login(ctx *gin.Context) {
	var req domain.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := authHandler.userService.Login(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie(
		"token", // name cookie
		token,   // value cookie
		authHandler.env.JWTExpiredHours*60*60, // expired from second
		"/",   // appicable path
		"",    // domain, (for localhost is blank)
		false, // secure  (set true when using ssl)
		true,  // http only (javascript cannot read this cookie)
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login is success",
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
