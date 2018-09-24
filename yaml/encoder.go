package yaml

import (
	"io"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/mickep76/encoding"
)

type encoder struct {
	encoder *yaml.Encoder
}

func (y *yamlEncoding) NewEncoder(w io.Writer, opts ...encoding.EncoderOption) (encoding.Encoder, error) {
	return &encoder{encoder: yaml.NewEncoder(w)}, nil
}

func (e *encoder) SetIndent(indent string) error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "encoding json")
}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
