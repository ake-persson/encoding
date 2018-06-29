package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

type decoder struct {
	reader io.Reader
}

func (d *decoder) SetMapString() {}

func (d *decoder) Decode(value interface{}) error {
	_, err := toml.DecodeReader(d.reader, value)
	return err
}
