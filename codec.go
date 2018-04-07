package encdec

import "io"

var codecs = make(map[string]Codec)

// Codec interface.
type Codec interface {
	SetIndent(indent string) error
	Encoder(writer io.Writer) (Encoder, error)
	Decoder(reader io.Reader) (Decoder, error)
}

// Register codec.
func Register(name string, codec Codec) {
	codecs[name] = codec
}
