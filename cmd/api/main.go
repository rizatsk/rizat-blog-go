package main

import (
	"fmt"
	"log"
	"service_blog_go/internal/config"
	"service_blog_go/internal/handler"
	"service_blog_go/internal/service"
	"service_blog_go/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadEnvironment()
	port := env.Port

	// connection db postgres
	db_postgres := database.InitPostgres(env)
	_ = db_postgres

	healthService := service.NewHealthService(env)
	healthHandler := handler.NewHealthHandler(healthService)

	route := gin.Default()

	v1 := route.Group("/api/v1")
	{
		v1.GET("/health", healthHandler.Ping)
	}

	serverAddr := fmt.Sprintf(":%s", port)
	if err := route.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
