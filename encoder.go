package encdec

import "io"

// Encoder interface.
type Encoder interface {
	SetIndent(indent string) error
	Encode(value interface{}) error
}

// NewEncoder constructor.
func NewEncoder(codec string, writer io.Writer, options ...func(Encoder) error) (Encoder, error) {
	c, ok := codecs[codec]
	if !ok {
		return nil, ErrUnknownCodec
	}

	enc := c.NewEncoder(writer)
	for _, option := range options {
		if err := option(enc); err != nil {
			return nil, err
		}
	}

	return enc, nil
}

// Indent output setter.
func Indent(indent string) func(Encoder) error {
	return func(e Encoder) error {
		return e.SetIndent(indent)
	}
}

// ToByte method.
func ToByte(codec string, value interface{}, options ...func(Encoder) error) ([]byte, error) {
	return nil, nil
}

// ToFile method.
func ToFile(codec string, file string, value interface{}, options ...func(Encoder) error) error {
	return nil
}
