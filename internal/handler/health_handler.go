package handler

import (
	"net/http"
	"service_blog_go/internal/domain"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	healthService domain.HealthService
}

func NewHealthHandler(service domain.HealthService) *HealthHandler {
	return &HealthHandler{healthService: service}
}

func (health *HealthHandler) Ping(ctx *gin.Context) {
	reponse := health.healthService.CheckHealth(ctx.Request.Context())
	ctx.JSON(http.StatusOK, reponse)
}
