package bson

import (
	"io"

	"github.com/mickep76/encdec"
	"github.com/pkg/bson"
)

type encoding struct{}

func (c *encoding) NewEncoder(writer io.Writer) encdec.Encoder {
	return &encoder{encoder: bson.NewEncoder(writer)}
}

func (c *encoding) NewDecoder(reader io.Reader) encdec.Decoder {
	return &decoder{decoder: bson.NewDecoder(reader)}
}

func init() {
	encdec.Register("bson", &encoding{})
}
