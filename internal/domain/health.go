package domain

import "context"

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type HealthService interface {
	CheckHealth(ctx context.Context) *HealthResponse
}
