package yaml

import "github.com/mickep76/encoding"

type yamlEncoding struct{}

func init() {
	encoding.Register("yaml", &yamlEncoding{})
}
