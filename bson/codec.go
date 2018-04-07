package bson

import (
	"io"

	"github.com/mickep76/encdec"
	"github.com/pkg/bson"
)

type codec struct{}

func (c *codec) NewEncoder(writer io.Writer) encdec.Encoder {
	return &encoder{encoder: bson.NewEncoder(writer)}
}

func (c *codec) NewDecoder(reader io.Reader) encdec.Decoder {
	return &decoder{decoder: bson.NewDecoder(reader)}
}

func init() {
	encdec.Register("bson", &codec{})
}
