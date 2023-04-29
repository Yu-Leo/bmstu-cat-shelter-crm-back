package apperror

var (
	ValidationError = NewAppError("Validation error")
	CatNotFound     = NewAppError("Cat not found")
)

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
func NewAppError(message string) *AppError {
	return &AppError{
		Message: message,
	}
}
