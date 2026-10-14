package service

import (
	"context"
	"service_blog_go/internal/config"
	"service_blog_go/internal/domain"
)

type healthService struct {
	env *config.Environment
}

func NewHealthService(env *config.Environment) domain.HealthService {
	return &healthService{
		env: env,
	}
}

func (service *healthService) CheckHealth(ctx context.Context) *domain.DataHealthResponse {
	return &domain.DataHealthResponse{
		AppName: service.env.AppName,
		AppEnv:  service.env.AppEnv,
	}
}
