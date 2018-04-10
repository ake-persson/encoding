package json

import (
	"encoding/json"

	"github.com/mickep76/encdec"
)

type decoder struct {
	decoder *json.Decoder
}

func (d *decoder) SetMapString() error {
	return encdec.ErrNotSupported
}

func (d *decoder) Decode(value interface{}) error {
	return d.decoder.Decode(value)
}
