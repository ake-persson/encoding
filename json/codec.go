package json

import (
	"encoding/json"
	"io"

	"github.com/mickep76/encdec"
)

type codec struct{}

func (c *codec) SetIndent(indent string) error {
	//	c.encoder.SetIndent("", indent)
	return nil
}

func (c *codec) Encoder(writer io.Writer) (encdec.Encoder, error) {
	return &encoder{
		json.NewEncoder(writer),
	}, nil
}

func (c *codec) Decoder(reader io.Reader) (encdec.Decoder, error) {
	return &decoder{
		json.NewDecoder(reader),
	}, nil
}

func init() {
	encdec.Register("json", &codec{})
}
