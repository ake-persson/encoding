package encoding

import (
	"bytes"
	"io"
)

// Decoder interface.
type Decoder interface {
	Decode(v interface{}) error
	SetMapString() error
}

// DecoderOption function.
type DecoderOption func(Decoder) error

// NewDecoder variadic constructor.
func NewDecoder(name string, r io.Reader, opts ...DecoderOption) (Decoder, error) {
	e, err := Registered(name)
	if err != nil {
		return nil, err
	}
	return e.NewDecoder(r, opts...)
}

// WithMapString convert map[interface{}]interface{} to map[string]interface{}.
// Suppored by YAML.
func WithMapString() DecoderOption {
	return func(d Decoder) error {
		return d.SetMapString()
	}
}

// Decode method.
func Decode(name string, encoded []byte, v interface{}, opts ...DecoderOption) error {
	d, err := NewDecoder(name, bytes.NewBuffer(encoded), opts...)
	if err != nil {
		return err
	}
	return d.Decode(v)
}
