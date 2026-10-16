package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Error     any    `json:"error,omitempty"`
	CreatedAt string `json:"created_at"`
}

func Success(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, ApiResponse{
		Status:    "success",
		Message:   message,
		Data:      data,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
}

func Failed(ctx *gin.Context, code int, message string, err any) {
	ctx.JSON(code, ApiResponse{
		Status:    "failed",
		Message:   message,
		Error:     err,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})
}
