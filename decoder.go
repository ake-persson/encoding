package encdec

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// Decoder interface.
type Decoder interface {
	Decode(value interface{}) error
}

// Decoder constructor.
func NewDecoder(encoding string, reader io.Reader, options ...func(Decoder) error) (Decoder, error) {
	c, ok := encodings[encoding]
	if !ok {
		return nil, ErrUnknownEncoding
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
func FromBytes(encoding string, encoded []byte, value interface{}, options ...func(Encoder) error) error {
	r := bufio.NewReader(bytes.NewReader(encoded))
	dec, err := NewDecoder(encoding, r)
	if err != nil {
		return err
	}

	if err := dec.Decode(value); err != nil {
		return err
	}

	return nil
}

// FromFile method.
func FromFile(encoding string, file string, value interface{}, options ...func(Encoder) error) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fp.Close()

	r := bufio.NewReader(fp)
	dec, err := NewDecoder(encoding, r)
	if err != nil {
		return err
	}

	if err := dec.Decode(value); err != nil {
		return err
	}

	return nil
}
