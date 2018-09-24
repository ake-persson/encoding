package yaml

import (
	"io"

	"gopkg.in/yaml.v2"

	"github.com/mickep76/encoding"
)

type decoder struct {
	decoder   *yaml.Decoder
	mapString bool
}

func (y *yamlEncoding) NewDecoder(reader io.Reader, opts ...encoding.DecoderOption) (encoding.Decoder, error) {
	d := &decoder{decoder: yaml.NewDecoder(reader)}
	for _, opt := range opts {
		if err := opt(d); err != nil {
			return nil, err
		}
	}
	return d, nil
}

func (d *decoder) SetMapString() error {
	d.mapString = true
	return nil
}

func (d *decoder) Decode(value interface{}) error {
	if d.mapString {
		var res interface{}
		if err := d.decoder.Decode(&res); err != nil {
			return err
		}

		*value.(*interface{}) = cleanupMapValue(res)
		return nil
	}

	return d.decoder.Decode(value)
}
