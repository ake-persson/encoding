package yaml

import (
	"gopkg.in/yaml.v2"

	"github.com/mickep76/encdec"
)

type encoder struct {
	encoder *yaml.Encoder
}

func (e *encoder) SetIndent(indent string) error {
	return encdec.ErrNotSupported
}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
