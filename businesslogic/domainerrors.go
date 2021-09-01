package businesslogic

import "fmt"

const (
	InvalidFormatKey = "client_error_invalid_format"
)

type DomainError struct {
	Message string
	Key     string
}

func NewInvalidFormatDomainError(message string) DomainError {
	return DomainError{
		Key:     InvalidFormatKey,
		Message: message,
	}
}

func (e DomainError) Error() string {

	return fmt.Sprintf("%s: %s", e.Key, e.Message)
}
