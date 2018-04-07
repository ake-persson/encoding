package json

import "encoding/json"

type decoder struct {
	*json.Decoder
}

func (d *decoder) Decode(bytes []byte, value interface{}) error {
	return nil
}
