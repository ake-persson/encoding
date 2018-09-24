package toml

import (
	"io"

	"github.com/BurntSushi/toml"

	"github.com/mickep76/encoding"
)

type encoding struct{}

func (c *encoding) NewEncoder(writer io.Writer) encoding.Encoder {
	return &encoder{encoder: toml.NewEncoder(writer)}
}

func (c *encoding) NewDecoder(reader io.Reader) encoding.Decoder {
	return &decoder{reader: reader}
}

func init() {
	encoding.Register("toml", &encoding{})
}
