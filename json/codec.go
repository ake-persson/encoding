package json

import (
	"encoding/json"
	"io"

	"github.com/mickep76/encdec"
)

type codec struct{}

func (c *codec) NewEncoder(writer io.Writer) (encdec.Encoder, error) {
	return &encoder{
		json.NewEncoder(writer),
	}, nil
}

func (c *codec) NewDecoder(reader io.Reader) (encdec.Decoder, error) {
	return &decoder{
		json.NewDecoder(reader),
	}, nil
}

func init() {
	encdec.Register("json", &codec{})
}
