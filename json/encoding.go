package json

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/mickep76/encoding"
)

type jsonEncoding struct {
	indent string
}

type jsonEncoder struct {
	encoder *json.Encoder
}

type jsonDecoder struct {
	decoder *json.Decoder
}

func (e *jsonEncoding) SetIndent(indent string) error {
	e.indent = indent
	return nil
}

func (e *jsonEncoding) SetMapString() error {
	return errors.Wrap(encoding.ErrUnsupportedOption, "algorithm gzip")
}

func (e *jsonEncoding) NewEncoder(w io.Writer) (encoding.Encoder, error) {
	return &jsonEncoder{encoder: json.NewEncoder(w)}, nil
}

func (e *jsonEncoding) Encode(v interface{}) ([]byte, error) {
	return encoding.Encode(e, v)
}

func (e *jsonEncoder) Encode(v interface{}) error {
	return e.encoder.Encode(v)
}

func (e *jsonEncoding) NewDecoder(r io.Reader) (encoding.Decoder, error) {
	return &jsonDecoder{decoder: json.NewDecoder(r)}, nil
}

func (e *jsonEncoding) Decode(b []byte, v interface{}) error {
	return encoding.Decode(e, b, v)
}

func (d *jsonDecoder) Decode(v interface{}) error {
	return d.decoder.Decode(v)
}

func init() {
	encoding.Register("json", &jsonEncoding{})
}
