package json

import (
	"encoding/json"
	"io"

	"github.com/mickep76/encdec"
)

type codec struct{}

func (c *codec) NewEncoder(writer io.Writer) encdec.Encoder {
	return &encoder{encoder: json.NewEncoder(writer)}
}

func (c *codec) NewDecoder(reader io.Reader) encdec.Decoder {
	return &decoder{decoder: json.NewDecoder(reader)}
}

func init() {
	encdec.Register("json", &codec{})
}
