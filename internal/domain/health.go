package domain

import "context"

type DataHealthResponse struct {
	AppName string `json:"appName"`
	AppEnv  string `json:"appEnv"`
}

type HealthResponse struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    DataHealthResponse `json:"data"`
}

type HealthService interface {
	CheckHealth(ctx context.Context) *HealthResponse
}
