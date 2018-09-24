package yaml

import (
	"io"

	"github.com/mickep76/encoding"

	"gopkg.in/yaml.v2"
)

type encoding struct{}

func (c *encoding) NewEncoder(writer io.Writer) encoding.Encoder {
	return &encoder{encoder: yaml.NewEncoder(writer)}
}

func (c *encoding) NewDecoder(reader io.Reader) encoding.Decoder {
	return &decoder{decoder: yaml.NewDecoder(reader)}
}

func init() {
	encoding.Register("yaml", &encoding{})
}
