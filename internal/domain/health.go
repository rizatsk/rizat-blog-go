package domain

import "context"

type DataHealthResponse struct {
	AppName string `json:"appName"`
	AppEnv  string `json:"appEnv"`
}

type HealthService interface {
	CheckHealth(ctx context.Context) *DataHealthResponse
}
