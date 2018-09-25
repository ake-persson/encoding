package yaml

import (
	"io"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/mickep76/encoding"
)

type yamlCodec struct {
	mapString bool
}

type yamlEncoder struct {
	encoder *yaml.Encoder
}

type yamlDecoder struct {
	decoder   *yaml.Decoder
	mapString bool
}

func (c *yamlCodec) NewCodec() encoding.Codec {
	return &yamlCodec{}
}

func (c *yamlCodec) SetIndent(indent string) error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "codec yaml")
}

func (c *yamlCodec) SetMapString() error {
	c.mapString = true
	return nil
}

func (c *yamlCodec) NewEncoder(w io.Writer) (encoding.Encoder, error) {
	return &yamlEncoder{encoder: yaml.NewEncoder(w)}, nil
}

func (c *yamlCodec) Encode(v interface{}) ([]byte, error) {
	return encoding.Encode(c, v)
}

func (e *yamlEncoder) Encode(v interface{}) error {
	return e.encoder.Encode(v)
}

func (c *yamlCodec) NewDecoder(r io.Reader) (encoding.Decoder, error) {
	return &yamlDecoder{decoder: yaml.NewDecoder(r)}, nil
}

func (c *yamlCodec) Decode(b []byte, v interface{}) error {
	return encoding.Decode(c, b, v)
}

func (d *yamlDecoder) Decode(v interface{}) error {
	if d.mapString {
		var res interface{}
		if err := d.decoder.Decode(&res); err != nil {
			return err
		}

		*v.(*interface{}) = cleanupMapValue(res)
		return nil
	}

	return d.decoder.Decode(v)
}

func init() {
	encoding.Register("yaml", &yamlCodec{})
}
