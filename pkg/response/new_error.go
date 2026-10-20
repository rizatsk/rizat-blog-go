package response

import "fmt"

type AppError struct {
	Code        int
	Message     string
	ResultError any
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

func NewAppError(code int, message string, err any) *AppError {
	if message == "" {
		message = "Internal server error"
	}
	return &AppError{
		Code:        code,
		Message:     message,
		ResultError: err,
	}
}
