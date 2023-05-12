package errors

var (
	CatChipNumberAlreadyExists = NewAppError("Cat's chip number already exists")
	PersonPhoneAlreadyExists   = NewAppError("Person's phone already exists")
	RoomNumberAlreadyExists    = NewAppError("Room's number already exists")
	ResidentNotFound           = NewAppError("Resident not found")
	GuardianNotFound           = NewAppError("Guardian not found")
	RoomNotFound               = NewAppError("Room not found")
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
