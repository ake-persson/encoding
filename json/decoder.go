package json

import (
	"encoding/json"
	"io"
)

type decoder struct {
	reader  io.Reader
	decoder *json.Decoder
}

func (d *decoder) Decode(value interface{}) error {
	return nil
}
