package yaml

import (
	"io"

	"github.com/mickep76/encdec"

	"gopkg.in/yaml.v2"
)

type codec struct{}

func (c *codec) NewEncoder(writer io.Writer) encdec.Encoder {
	return &encoder{encoder: yaml.NewEncoder(writer)}
}

func (c *codec) NewDecoder(reader io.Reader) encdec.Decoder {
	return &decoder{decoder: yaml.NewDecoder(reader)}
}

func init() {
	encdec.Register("yaml", &codec{})
}
