package json

import (
	"encoding/json"
)

type encoder struct {
	*json.Encoder
}

func (e *encoder) SetIndent(indent string) {
	e.Encoder.SetIndent("", indent)
}

func (e *encoder) Encode(value interface{}) error {
	return e.Encode(value)
}
