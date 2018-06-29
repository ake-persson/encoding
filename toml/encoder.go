package toml

import (
	"github.com/BurntSushi/toml"
)

type encoder struct {
	encoder *toml.Encoder
}

func (e *encoder) SetIndent(indent string) {}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
