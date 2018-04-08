package json

import "encoding/json"

type encoder struct {
	encoder *json.Encoder
}

func (e *encoder) SetIndent(indent string) error {
	e.encoder.SetIndent("", indent)
	return nil
}

func (e *encoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}
