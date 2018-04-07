package json

import (
	"encoding/json"
)

type encoder struct {
	*json.Encoder
}

func (e *encoder) Encode(value interface{}) error {
	return e.Encode(value)
}
