package json

import "github.com/mickep76/encoding"

type jsonEncoding struct{}

func init() {
	encoding.Register("json", &jsonEncoding{})
}
