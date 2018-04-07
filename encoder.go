package encdec

import "io"

// Encoder interface.
type Encoder interface {
	Encode(value interface{}) error
}

// NewEncoder constructor.
func NewEncoder(codec string, writer io.Writer, options ...func(Codec) error) (Encoder, error) {
	c, ok := codecs[codec]
	if !ok {
		return nil, ErrUnknownCodec
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c.Encoder()
}

// Indent output.
func Indent(indent string) func(Codec) error {
	return func(c Codec) error {
		return c.SetIndent(indent)
	}
}
