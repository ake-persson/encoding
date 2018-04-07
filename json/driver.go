package json

import (
	"encoding/json"

	"github.com/mickep76/encdec"
)

type driver struct {
	encoder json.Encoder
	decoder json.Decoder
}

func (d *driver) SetIndent(indent string) error {
	d.encoder.SetIndent("", indent)
	return nil
}

func (d *driver) Encoder() (encdec.Encoder, error) {
	return &encoder{&d.encoder}, nil
}

func (d *driver) Decoder() (encdec.Decoder, error) {
	return &decoder{&d.decoder}, nil
}

func init() {
	encdec.Register("json", &driver{})
}
