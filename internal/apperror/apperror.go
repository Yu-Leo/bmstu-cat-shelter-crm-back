package apperror

var (
	CatChipNumberAlreadyExists = NewAppError("Cat's chip number already exists")
	PersonPhoneAlreadyExists   = NewAppError("Person's phone already exists")
	CatNotFound                = NewAppError("Cat not found")
	ResidentNotFound           = NewAppError("Resident not found")
	GuardianNotFound           = NewAppError("Guardian not found")
)

var (
	InternalServerError = ErrorJSON{Message: InternalServerErrorMsg}
)

const (
	ValidationErrorMsg     = "Validation error"
	InternalServerErrorMsg = "Internal Server Error"
	InvalidGuardianIdMsg   = "Invalid guardian id"
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
