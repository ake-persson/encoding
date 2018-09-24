package toml

import "github.com/mickep76/encoding"

type tomlEncoding struct{}

func init() {
	encoding.Register("toml", &tomlEncoding{})
}
