package encdec

import "errors"

var (
	ErrNotSupported = errors.New("method is not supported by codec")
	ErrUnknownCodec = errors.New("unknown codec")
)
