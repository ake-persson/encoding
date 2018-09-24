package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"

	"github.com/mickep76/encoding"
)

type decoder struct {
	decoder io.Reader
}

func (t *tomlEncoding) NewDecoder(r io.Reader, opts ...encoding.DecoderOption) (encoding.Decoder, error) {
	return &decoder{decoder: r}, nil
}

func (d *decoder) SetMapString() error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "encoding toml")
}

func (d *decoder) Decode(v interface{}) error {
	_, err := toml.DecodeReader(d.decoder, v)
	return err
}
