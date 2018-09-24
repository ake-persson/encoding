package encoding

import (
	"bytes"
	"io"
)

// Encoder interface.
type Encoder interface {
	Encode(v interface{}) error
	SetIndent(indent string)
}

// EncoderOption variadic function.
type EncoderOption func(Encoder)

// NewEncoder variadic constructor.
func NewEncoder(name string, w io.Writer, opts ...EncoderOption) (Encoder, error) {
	e, err := Registered(name)
	if err != nil {
		return nil, err
	}
	return e.NewEncoder(w, opts...)
}

// WithIndent output setter for JSON.
func WithIndent(indent string) EncoderOption {
	return func(e Encoder) {
		e.SetIndent(indent)
	}
}

// Encode method.
func Encode(name string, v interface{}, opts ...EncoderOption) ([]byte, error) {
	var buf bytes.Buffer

	e, err := NewEncoder(name, &buf, opts...)
	if err != nil {
		return nil, err
	}

	if err := e.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
