package util

type AppError struct {
	Message string `json:"message"`
}

func NewAppError(err error) *AppError {
	return &AppError{
		Message: err.Error(),
	}
}
