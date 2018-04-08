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
func NewDecoder(codec string, reader io.Reader, options ...func(Decoder) error) (Decoder, error) {
	c, ok := codecs[codec]
	if !ok {
		return nil, ErrUnknownCodec
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
func FromBytes(codec string, encoded []byte, value interface{}, options ...func(Encoder) error) error {
	r := bufio.NewReader(bytes.NewReader(encoded))
	dec, err := NewDecoder(codec, r)
	if err != nil {
		return err
	}

	if err := dec.Decode(value); err != nil {
		return err
	}

	return nil
}

// FromFile method.
func FromFile(codec string, file string, value interface{}, options ...func(Encoder) error) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fp.Close()

	r := bufio.NewReader(fp)
	dec, err := NewDecoder(codec, r)
	if err != nil {
		return err
	}

	if err := dec.Decode(value); err != nil {
		return err
	}

	return nil
}
