package handler

import (
	"net/http"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"
	"service_blog_go/pkg/response"

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
		formatedErr := response.FormValidationError(err)
		response.Failed(
			ctx,
			http.StatusUnprocessableEntity,
			"Bad request",
			formatedErr,
		)

		return
	}

	token, err := authHandler.userService.Login(ctx.Request.Context(), &req)
	if err != nil {
		response.Failed(
			ctx,
			http.StatusUnauthorized,
			"Failed login",
			err.Error(),
		)
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

	response.Success(
		ctx,
		http.StatusOK,
		"Login success",
		nil,
	)
}
