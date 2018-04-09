package encdec

import "errors"

var (
	ErrNotSupported  = errors.New("method is not supported by encoding")
	ErrNotRegistered = errors.New("encoding isn't registered")
)
