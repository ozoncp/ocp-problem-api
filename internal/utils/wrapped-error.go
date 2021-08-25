package utils

import (
	"errors"
	"strings"
)

type wrappedError struct {
	err error
	message string
}

func (werr *wrappedError) Error() string {
	builder := &strings.Builder{}
	builder.WriteString(werr.message)

	currentError := werr.err
	for ; currentError != nil; {
		builder.WriteString(", ")
		builder.WriteString(currentError.Error())
		currentError = errors.Unwrap(currentError)
	}

	return builder.String()
}

func (werr *wrappedError) Unwrap() error {
	return werr.err
}

func NewWrappedError(message string, err error) error {
	return &wrappedError{
		message: message,
		err: err,
	}
}