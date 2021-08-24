package utils

type wrappedError struct {
	err error
	message string
}

func (werr *wrappedError) Error() string {
	return werr.message
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