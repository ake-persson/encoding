package json

import (
	"testing"

	"github.com/mickep76/encoding"
)

func TestEncode(t *testing.T) {
	b, err := encoding.Encode("json", testMap)
	if err != nil {
		t.Error(err)
	}

	if string(b) != testEncoded {
		t.Errorf("want:\n%s, got:\n%s", testEncoded, string(b))
	}
}

func TestEncodeWithIndent(t *testing.T) {
	b, err := encoding.Encode("json", testMap, encoding.WithIndent("  "))
	if err != nil {
		t.Error(err)
	}

	if string(b) != testEncodedIndent {
		t.Errorf("want:\n%s, got:\n%s", testEncodedIndent, string(b))
	}
}
