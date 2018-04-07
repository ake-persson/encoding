package json

import (
	"encoding/json"

	"github.com/mickep76/encdec"
)

type codec struct {
	encoder json.Encoder
	decoder json.Decoder
}

func (c *codec) SetIndent(indent string) error {
	c.encoder.SetIndent("", indent)
	return nil
}

func (c *codec) Encoder() (encdec.Encoder, error) {
	return &encoder{encoder: &c.encoder}, nil
}

func (c *codec) Decoder() (encdec.Decoder, error) {
	return &decoder{decoder: &c.decoder}, nil
}

func init() {
	encdec.Register("json", &codec{})
}
