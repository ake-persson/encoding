package encdec

import "io"

// Decoder interface.
type Decoder interface {
	Decode(value interface{}) error
}

// Decoder constructor.
func NewDecoder(codec string, reader io.Reader, options ...func(Codec) error) (Decoder, error) {
	c, ok := codecs[codec]
	if !ok {
		return nil, ErrUnknownCodec
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c.Decoder()
}
