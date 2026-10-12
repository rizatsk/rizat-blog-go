package service

import (
	"context"
	"service_blog_go/internal/domain"
)

type healthService struct{}

func NewHealthService() domain.HealthService {
	return &healthService{}
}

func (s *healthService) CheckHealth(ctx context.Context) *domain.HealthResponse {
	return &domain.HealthResponse{
		Status:  "OK",
		Message: "Server is healthy and running smoothly",
	}
}
