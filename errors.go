package encdec

import "errors"

var (
	ErrNotSupported = errors.New("method is not supported by encoding")
	ErrUnknownEncoding = errors.New("unknown encoding")
)
