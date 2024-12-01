package helpers

type CustomError struct {
	Description string
	Code int
}

func NewError(errorMessage string, errorCode int) *CustomError {
	return &CustomError{Description: errorMessage, Code: errorCode}
}