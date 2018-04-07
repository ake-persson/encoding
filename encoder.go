package encdec

import "io"

// Encoder interface.
type Encoder interface {
	SetIndent(indent string)
	Encode(value interface{}) error
}

// NewEncoder constructor.
func NewEncoder(codec string, writer io.Writer, options ...func(Encoder) error) (Encoder, error) {
	c, ok := codecs[codec]
	if !ok {
		return nil, ErrUnknownCodec
	}

	enc, err := c.NewEncoder(writer)
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		if err := option(enc); err != nil {
			return nil, err
		}
	}

	return enc, nil
}

// Indent output.
func Indent(indent string) func(Encoder) error {
	return func(e Encoder) error {
		e.SetIndent(indent)
		return nil
	}
}
