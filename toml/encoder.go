package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	"github.com/mickep76/encoding"
)

type encoder struct {
	encoder *toml.Encoder
}

func (t *tomlEncoding) NewEncoder(w io.Writer, opts ...encoding.EncoderOption) (encoding.Encoder, error) {
	e := &encoder{encoder: toml.NewEncoder(w)}
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (e *encoder) SetIndent(indent string) error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "encoding toml")
}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
