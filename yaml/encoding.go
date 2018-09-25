package yaml

import (
	"io"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/mickep76/encoding"
)

type yamlEncoding struct {
	mapString bool
}

type yamlEncoder struct {
	encoder *yaml.Encoder
}

type yamlDecoder struct {
	decoder   *yaml.Decoder
	mapString bool
}

func (e *yamlEncoding) SetIndent(indent string) error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "algorithm yaml")
}

func (e *yamlEncoding) SetMapString() error {
	e.mapString = true
	return nil
}

func (e *yamlEncoding) NewEncoder(w io.Writer) (encoding.Encoder, error) {
	return &yamlEncoder{encoder: yaml.NewEncoder(w)}, nil
}

func (e *yamlEncoding) Encode(v interface{}) ([]byte, error) {
	return encoding.Encode(e, v)
}

func (e *yamlEncoder) Encode(v interface{}) error {
	return e.encoder.Encode(v)
}

func (e *yamlEncoding) NewDecoder(r io.Reader) (encoding.Decoder, error) {
	return &yamlDecoder{decoder: yaml.NewDecoder(r)}, nil
}

func (e *yamlEncoding) Decode(b []byte, v interface{}) error {
	return encoding.Decode(e, b, v)
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
	encoding.Register("yaml", &yamlEncoding{})
}
