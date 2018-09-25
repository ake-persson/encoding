package encoding

import (
	"bytes"
	"fmt"
	"io"
)

var encodings = make(map[string]Encoding)

// Encoding interface.
type Encoding interface {
	NewEncoding() Encoding
	NewEncoder(w io.Writer) (Encoder, error)
	NewDecoder(r io.Reader) (Decoder, error)
	Encode(v interface{}) ([]byte, error)
	Decode(b []byte, v interface{}) error
	SetIndent(indent string) error
	SetMapString() error
}

// Encoder interface.
type Encoder interface {
	Encode(v interface{}) error
}

// Decoder interface.
type Decoder interface {
	Decode(v interface{}) error
}

// Option variadic function.
type Option func(Encoding) error

// Register encoding.
func Register(name string, encoding Encoding) {
	encodings[name] = encoding
}

// Encodings registered.
func Encodings() []string {
	l := []string{}
	for a := range encodings {
		l = append(l, a)
	}
	return l
}

// NewEncoding variadic constructor.
func NewEncoding(name string, opts ...Option) (Encoding, error) {
	e, ok := encodings[name]
	if !ok {
		return nil, fmt.Errorf("encoding not registered: %s", name)
	}
	e = e.NewEncoding()
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}
	return e, nil
}

// WithIndent indent output.
// Supported by JSON.
func WithIndent(indent string) Option {
	return func(e Encoding) error {
		return e.SetIndent(indent)
	}
}

// WithMapString convert map[interface{}]interface{} to map[string]interface{}.
// Suppored by YAML.
func WithMapString() Option {
	return func(e Encoding) error {
		return e.SetMapString()
	}
}

// Encode method.
func Encode(e Encoding, v interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc, err := e.NewEncoder(&buf)
	if err != nil {
		return nil, err
	}

	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decode method.
func Decode(e Encoding, encoded []byte, v interface{}) error {
	d, err := e.NewDecoder(bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}
	return d.Decode(v)
}
