package yaml

import "gopkg.in/yaml.v2"

type decoder struct {
	mapString bool
	decoder   *yaml.Decoder
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
	} else {
		if err := d.decoder.Decode(value); err != nil {
			return err
		}
	}

	return nil
}
