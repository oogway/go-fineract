package fineract

import "encoding/json"

type FineractError struct {
	Code    string
	Message *json.RawMessage
}

func (f *FineractError) Error() string {
	return f.Code
}

const (
	// ErrCodeSerialization is the serialization error code that is received
	// during protocol unmarshaling.
	ErrCodeSerialization = "SerializationError"

	// ErrCodeRead is an error that is returned during HTTP when resource not found.
	ErrNotFound = "ResourceNotFound"

	// ErrBadRequest is an error that is returned during HTTP when when a request cannot be processed.
	ErrBadRequest = "StatusBadRequest"

	// ErrCodeResponseTimeout is the connection timeout error that is received
	// during body reads.
	ErrCodeResponseTimeout = "ResponseTimeout"

	// ErrInternalServer is the internal server error
	ErrInternalServer = "InternalServerError"

	// ErrAuthenticationFailure is the code returned when there is a athentication failure
	ErrAuthenticationFailure = "AuthenticationFailure"
)

func GetFineractStatusCode(code int) string {
	switch code {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrAuthenticationFailure
	case 404:
		return ErrNotFound
	case 500:
		return ErrInternalServer
	}
	return ErrInternalServer
}
