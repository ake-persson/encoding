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
