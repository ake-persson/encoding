package json

import (
	"encoding/json"
	"io"

	"github.com/mickep76/encoding"
)

type encoding struct{}

func (c *encoding) NewEncoder(writer io.Writer) encoding.Encoder {
	return &encoder{encoder: json.NewEncoder(writer)}
}

func (c *encoding) NewDecoder(reader io.Reader) encoding.Decoder {
	return &decoder{decoder: json.NewDecoder(reader)}
}

func init() {
	encoding.Register("json", &encoding{})
}
