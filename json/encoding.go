package json

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/encoding"
)

type jsonCodec struct {
	indent string
}

type jsonEncoder struct {
	encoder *json.Encoder
}

type jsonDecoder struct {
	decoder *json.Decoder
}

func (c *jsonCodec) NewCodec() encoding.Codec {
	return &jsonCodec{}
}

func (c *jsonCodec) SetIndent(indent string) error {
	c.indent = indent
	return nil
}

func (c *jsonCodec) SetMapString() error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "codec json")
}

func (c *jsonCodec) NewEncoder(w io.Writer) (encoding.Encoder, error) {
	e := &jsonEncoder{encoder: json.NewEncoder(w)}
	if c.indent != "" {
		e.encoder.SetIndent("", c.indent)
	}
	return e, nil
}

func (c *jsonCodec) Encode(v interface{}) ([]byte, error) {
	return encoding.Encode(c, v)
}

func (e *jsonEncoder) Encode(v interface{}) error {
	return e.encoder.Encode(v)
}

func (c *jsonCodec) NewDecoder(r io.Reader) (encoding.Decoder, error) {
	return &jsonDecoder{decoder: json.NewDecoder(r)}, nil
}

func (c *jsonCodec) Decode(b []byte, v interface{}) error {
	return encoding.Decode(c, b, v)
}

func (d *jsonDecoder) Decode(v interface{}) error {
	return d.decoder.Decode(v)
}

func init() {
	encoding.Register("json", &jsonCodec{})
}
