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

func (h *HealthHandler) Ping(c *gin.Context) {
	reponse := h.healthService.CheckHealth(c.Request.Context())
	c.JSON(http.StatusOK, reponse)
}
