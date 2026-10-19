package middleware

import (
	"log"
	"net/http"
	"service_blog_go/internal/config"
	"service_blog_go/pkg/tokenize"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(env *config.Environment) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			log.Println("Token is not found from cookie")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			ctx.Abort()
			return
		}

		claims := &tokenize.JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			return []byte(env.JWTSecretKey), nil
		})
		if err != nil || !token.Valid {
			log.Println(ctx, "Token is not valid")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("name", claims.Name)
		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Set("created_at", claims.CreatedAt)
		ctx.Set("updated_at", claims.UpdatedAt)
		ctx.Next()
	}
}
