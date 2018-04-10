package encdec

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Encoder interface.
type Encoder interface {
	SetIndent(indent string) error
	Encode(value interface{}) error
}

// NewEncoder constructor.
func NewEncoder(encoding string, writer io.Writer, options ...func(Encoder) error) (Encoder, error) {
	c, ok := encodings[encoding]
	if !ok {
		return nil, ErrNotRegistered
	}

	enc := c.NewEncoder(writer)
	for _, option := range options {
		if err := option(enc); err != nil {
			return nil, err
		}
	}

	return enc, nil
}

// WithIndent output setter.
func WithIndent(indent string) func(Encoder) error {
	return func(e Encoder) error {
		return e.SetIndent(indent)
	}
}

// ToBytes method.
func ToBytes(encoding string, value interface{}, options ...func(Encoder) error) ([]byte, error) {
	var buf bytes.Buffer

	enc, err := NewEncoder(encoding, &buf, options...)
	if err != nil {
		return nil, err
	}

	if err := enc.Encode(value); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ToFile method.
func ToFile(encoding string, file string, value interface{}, options ...func(Encoder) error) error {
	fp, err := os.Create(file)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(fp)
	enc, err := NewEncoder(encoding, w, options...)
	if err != nil {
		return err
	}

	if err := enc.Encode(value); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return fp.Close()
}
