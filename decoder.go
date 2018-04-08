package encdec

import "io"

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

// FromByte method.
func FromByte(codec string, encoded []byte, value interface{}, options ...func(Encoder) error) error {
	return nil
}

// FromFile method.
func FromFile(codec string, file string, value interface{}, options ...func(Encoder) error) error {
	return nil
}
