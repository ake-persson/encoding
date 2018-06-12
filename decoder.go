package encdec

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Decoder interface.
type Decoder interface {
	SetMapString() error
	Decode(value interface{}) error
}

// DecoderOption function.
type DecoderOption func(Decoder) error

// WithMapString convert map[interface{}]interface{} to map[string]interface{} for YAML.
func WithMapString() DecoderOption {
	return func(d Decoder) error {
		return d.SetMapString()
	}
}

// NewDecoder variadic constructor.
func NewDecoder(encoding string, reader io.Reader, options ...DecoderOption) (Decoder, error) {
	c, ok := encodings[encoding]
	if !ok {
		return nil, ErrNotRegistered
	}

	dec := c.NewDecoder(reader)
	for _, option := range options {
		if err := option(dec); err != nil {
			return nil, err
		}
	}

	return dec, nil
}

// FromBytes method.
func FromBytes(encoding string, encoded []byte, value interface{}, options ...DecoderOption) error {
	r := bufio.NewReader(bytes.NewReader(encoded))
	dec, err := NewDecoder(encoding, r, options...)
	if err != nil {
		return err
	}

	return dec.Decode(value)
}

// FromFile method.
func FromFile(encoding string, file string, value interface{}, options ...DecoderOption) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}

	r := bufio.NewReader(fp)
	dec, err := NewDecoder(encoding, r, options...)
	if err != nil {
		return err
	}

	if err := dec.Decode(value); err != nil {
		return err
	}

	return fp.Close()
}
