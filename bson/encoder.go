package bson

import "github.com/pkg/bson"

type encoder struct {
	encoder *bson.Encoder
}

func (e *encoder) SetIndent(indent string) {
}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
