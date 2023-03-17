package apperror

const (
	ValidationErrorMsg     = "Validation error"
	InternalServerErrorMsg = "Internal Server Error"
)

type ErrorJSON struct {
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

type AppError struct {
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}
