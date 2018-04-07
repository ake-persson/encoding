package encdec

import "io"

var codecs = make(map[string]Codec)

// Codec interface.
type Codec interface {
	NewEncoder(writer io.Writer) Encoder
	NewDecoder(reader io.Reader) Decoder
}

// Register codec.
func Register(name string, codec Codec) {
	codecs[name] = codec
}
