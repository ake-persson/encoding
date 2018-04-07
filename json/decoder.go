package json

import (
	"encoding/json"
)

type decoder struct {
	*json.Decoder
}

func (d *decoder) Decode(value interface{}) error {
	return d.Decode(value)
}
