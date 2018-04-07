package json

import (
	"encoding/json"
	"io"

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

func (c *codec) Encoder(writer io.Writer) (encdec.Encoder, error) {
	return &encoder{encoder: &c.encoder, writer: writer}, nil
}

func (c *codec) Decoder(reader io.Reader) (encdec.Decoder, error) {
	return &decoder{decoder: &c.decoder, reader: reader}, nil
}

func init() {
	encdec.Register("json", &codec{})
}
