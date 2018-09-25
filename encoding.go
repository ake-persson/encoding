package encoding

import (
	"bytes"
	"fmt"
	"io"
)

var codecs = make(map[string]Codec)

// Codec interface.
type Codec interface {
	NewCodec() Codec
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
type Option func(Codec) error

// Register codec.
func Register(name string, codec Codec) {
	codecs[name] = codec
}

// Codecs registered.
func Codecs() []string {
	l := []string{}
	for a := range codecs {
		l = append(l, a)
	}
	return l
}

// Registered is the algorithm registered.
func Registered(name string) error {
	_, ok := codecs[name]
	if !ok {
		return fmt.Errorf("codec not registered: %s", name)
	}
	return nil
}

// NewCodec variadic constructor.
func NewCodec(name string, opts ...Option) (Codec, error) {
	e, ok := codecs[name]
	if !ok {
		return nil, fmt.Errorf("codec not registered: %s", name)
	}
	e = e.NewCodec()
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
	return func(c Codec) error {
		return c.SetIndent(indent)
	}
}

// WithMapString convert map[interface{}]interface{} to map[string]interface{}.
// Suppored by YAML.
func WithMapString() Option {
	return func(c Codec) error {
		return c.SetMapString()
	}
}

// Encode method.
func Encode(c Codec, v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	e, err := c.NewEncoder(&buf)
	if err != nil {
		return nil, err
	}
	if err := e.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode method.
func Decode(c Codec, encoded []byte, v interface{}) error {
	d, err := c.NewDecoder(bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}
	return d.Decode(v)
}
