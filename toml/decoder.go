package toml

import (
	"github.com/pelletier/go-toml"

	"github.com/mickep76/encdec"
)

type decoder struct {
	decoder *toml.Decoder
}

func (d *decoder) SetMapString() error {
	return encdec.ErrNotSupported
}

func (d *decoder) Decode(value interface{}) error {
	return d.decoder.Decode(value)
}
