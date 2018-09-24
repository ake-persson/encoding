package json

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/encoding"
)

type decoder struct {
	decoder *json.Decoder
}

func (j *jsonEncoding) NewDecoder(r io.Reader, opts ...encoding.DecoderOption) (encoding.Decoder, error) {
	d := &decoder{decoder: json.NewDecoder(r)}
	for _, opt := range opts {
		if err := opt(d); err != nil {
			return nil, err
		}
	}
	return d, nil
}

func (d *decoder) SetMapString() error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "encoding json")
}

func (d *decoder) Decode(v interface{}) error {
	return d.decoder.Decode(v)
}
