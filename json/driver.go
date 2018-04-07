package json

import (
	"github.com/mickep76/encdec"
)

type driver struct {
	indent string
}

func (d *driver) SetIndent(indent string) error {
	d.indent = indent
	return nil
}

func (d *driver) Encoder() (encdec.Encoder, error) {
	return &encoder{}, nil
}

func (d *driver) Decoder() (encdec.Decoder, error) {
	return &decoder{}, nil
}

func init() {
	encdec.Register("json", &driver{})
}
