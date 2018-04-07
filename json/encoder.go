package json

import "encoding/json"

type encoder struct {
	*json.Encoder
}

func (e *encoder) Encode(value interface{}) ([]byte, error) {
	return []byte(""), nil
}
