package handler

import (
	"net/http"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"
	"service_blog_go/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService domain.AuthService
	env         *config.Environment
}

func NewAuthHandler(authService domain.AuthService, env *config.Environment) *AuthHandler {
	return &AuthHandler{
		authService: authService,
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

	token, err := authHandler.authService.Login(ctx.Request.Context(), &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "Email or password is wrong" {
			statusCode = http.StatusUnauthorized
		}
		response.Failed(
			ctx,
			statusCode,
			"Failed login",
			err.Error(),
		)
		return
	}

	ctx.SetCookie(
		"token",           // name cookie
		token.TokenString, // value cookie
		authHandler.env.JWTExpiredHours*60*60, // expired from second
		"/",   // appicable path
		"",    // domain, (for localhost is blank)
		false, // secure  (set true when using ssl)
		true,  // http only (javascript cannot read this cookie)
	)

	ctx.SetCookie(
		"refreshToken",           // name cookie
		token.RefreshTokenString, // value cookie
		authHandler.env.JWTRfreshTokenExpiredDays*24*60*int(time.Hour), // expired from second
		"/api/v1/auth/refresh", // appicable path
		"",                     // domain, (for localhost is blank)
		false,                  // secure  (set true when using ssl)
		true,                   // http only (javascript cannot read this cookie)
	)

	response.Success(
		ctx,
		http.StatusOK,
		"Login success",
		nil,
	)
}
