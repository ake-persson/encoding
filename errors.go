package encdec

import "errors"

var (
	// ErrNotSupported method is not supported by encoding.
	ErrNotSupported = errors.New("method is not supported by encoding")

	// ErrNotRegistered encoding isn't registered.
	ErrNotRegistered = errors.New("encoding isn't registered")
)
