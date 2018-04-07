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

	dec, err := c.NewDecoder(reader)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		if err := option(dec); err != nil {
			return nil, err
		}
	}

	return dec, nil
}
