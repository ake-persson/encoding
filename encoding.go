package encoding

import (
	"fmt"
	"io"
)

var encodings = make(map[string]Encoding)

// Encoding interface.
type Encoding interface {
	NewEncoder(w io.Writer, opts ...EncoderOption) (Encoder, error)
	NewDecoder(r io.Reader, opts ...DecoderOption) (Decoder, error)
}

// Register encoding.
func Register(name string, encoding Encoding) {
	encodings[name] = encoding
}

// Registered encoding.
func Registered(name string) (Encoding, error) {
	e, ok := encodings[name]
	if !ok {
		return nil, fmt.Errorf("encoding not registered: %s", name)
	}
	return e, nil
}

// Encodings registered.
func Encodings() []string {
	l := []string{}
	for a := range encodings {
		l = append(l, a)
	}
	return l
}
