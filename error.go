package fineractor

import "encoding/json"

type FineractError struct {
	Code    string
	Message *json.RawMessage
}

func (f *FineractError) Error() string {
	return ""
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

	ErrInternalServer = "InternalServerError"
)

func GetFineractStatusCode(code int) string {
	switch code {
	case 404:
		return ErrNotFound
	case 400:
		return ErrBadRequest
	}
	return ErrInternalServer
}
