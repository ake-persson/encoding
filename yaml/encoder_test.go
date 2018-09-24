package yaml

import (
	"testing"

	"github.com/mickep76/encoding"
)

func TestEncoder(t *testing.T) {
	b, err := encoding.Encode("yaml", testMap)
	if err != nil {
		t.Error(err)
	}

	if string(b) != testEncoded {
		t.Errorf("want:\n%s, got:\n%s", testEncoded, string(b))
	}
}

func TestToByteWithIndent(t *testing.T) {
	if _, err := encoding.Encode("yaml", testMap, encoding.WithIndent("")); err != nil {
		t.Error(err)
	}
}
