package toml

import (
	"github.com/pelletier/go-toml"

	"github.com/mickep76/encdec"
)

type encoder struct {
	encoder *toml.Encoder
}

func (e *encoder) SetIndent(indent string) error {
	return encdec.ErrNotSupported
}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
