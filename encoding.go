package encdec

import "io"

var encodings = make(map[string]Encoding)

// Encoding interface.
type Encoding interface {
	NewEncoder(writer io.Writer) Encoder
	NewDecoder(reader io.Reader) Decoder
}

// Register encoding.
func Register(name string, encoding Encoding) {
	encodings[name] = encoding
}

// Registered encoding is registered.
func Registered(name string) error {
	_, ok := encodings[name]
	if !ok {
		return ErrNotRegistered
	}
	return nil
}

// Encodings registered.
func Encodings() []string {
	encs := []string{}
	for k, _ := range encodings {
		encs = append(encs, k)
	}
	return encs
}
