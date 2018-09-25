package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	"github.com/mickep76/encoding"
)

type tomlEncoding struct{}

type tomlEncoder struct {
	encoder *toml.Encoder
}

type tomlDecoder struct {
	decoder io.Reader
}

func (e *tomlEncoding) NewEncoding() encoding.Encoding {
	return &tomlEncoding{}
}

func (e *tomlEncoding) SetIndent(indent string) error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "algorithm toml")
}

func (e *tomlEncoding) SetMapString() error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "algorithm toml")
}

func (e *tomlEncoding) NewEncoder(w io.Writer) (encoding.Encoder, error) {
	return &tomlEncoder{encoder: toml.NewEncoder(w)}, nil
}

func (e *tomlEncoding) Encode(v interface{}) ([]byte, error) {
	return encoding.Encode(e, v)
}

func (e *tomlEncoder) Encode(value interface{}) error {
	return e.encoder.Encode(value)
}

func (e *tomlEncoding) NewDecoder(r io.Reader) (encoding.Decoder, error) {
	return &tomlDecoder{decoder: r}, nil
}

func (e *tomlEncoding) Decode(b []byte, v interface{}) error {
	return encoding.Decode(e, b, v)
}

func (d *tomlDecoder) Decode(v interface{}) error {
	_, err := toml.DecodeReader(d.decoder, v)
	return err
}

func init() {
	encoding.Register("toml", &tomlEncoding{})
}
