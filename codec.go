package encdec

var codecs = make(map[string]Codec)

// Codec interface.
type Codec interface {
	SetIndent(indent string) error
	Encoder() (Encoder, error)
	Decoder() (Decoder, error)
}

// Register codec.
func Register(name string, codec Codec) {
	codecs[name] = codec
}
