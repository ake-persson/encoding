package yaml

import "gopkg.in/yaml.v2"

type decoder struct {
	decoder *yaml.Decoder
}

func (d *decoder) Decode(value interface{}) error {
	return d.decoder.Decode(value)
}
