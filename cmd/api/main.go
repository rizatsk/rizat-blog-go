package main

import (
	"fmt"
	"log"
	"service_blog_go/internal/config"
	"service_blog_go/internal/handler"
	"service_blog_go/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	environemt := config.LoadEnvironment()
	port := environemt.Port

	healthService := service.NewHealthService()
	healthHandler := handler.NewHealthHandler(healthService)

	route := gin.Default()

	v1 := route.Group("/api/v1")
	{
		v1.GET("/health", healthHandler.Ping)
	}

	serverAddr := fmt.Sprintf(":%s", port)
	if err := route.Run(serverAddr); err != nil {
		log.Fatal("Failed to run server: %v", err)
	}
}
