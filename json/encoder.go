package json

import (
	"encoding/json"
	"io"
)

type encoder struct {
	writer  io.Writer
	encoder *json.Encoder
}

func (e *encoder) Encode(value interface{}) error {
	return nil
}
