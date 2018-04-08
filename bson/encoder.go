package bson

import (
	"github.com/pkg/bson"

	"github.com/mickep76/encdec"
)

type encoder struct {
	encoder *bson.Encoder
}

func (e *encoder) SetIndent(indent string) error {
	return encdec.ErrNotSupported
}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
