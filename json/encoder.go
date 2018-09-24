package json

import (
	"encoding/json"
	"io"

	"github.com/mickep76/encoding"
)

type encoder struct {
	encoder *json.Encoder
}

func (j *jsonEncoding) NewEncoder(w io.Writer, opts ...encoding.EncoderOption) (encoding.Encoder, error) {
	e := &encoder{encoder: json.NewEncoder(w)}
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}
	return e, nil
}

func (e *encoder) SetIndent(indent string) error {
	e.encoder.SetIndent("", indent)
	return nil
}

func (e *encoder) Encode(v interface{}) error {
	return e.encoder.Encode(v)
}
