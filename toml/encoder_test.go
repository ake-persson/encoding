package toml

import (
	"testing"

	"github.com/mickep76/encoding"
)

func TestEncode(t *testing.T) {
	b, err := encoding.Encode("toml", testStruct)
	if err != nil {
		t.Error(err)
	}

	if string(b) != testEncoded {
		t.Errorf("want:\n%s\n\ngot:\n%s", testEncoded, string(b))
	}
}

func TestEncodeWithIndent(t *testing.T) {
	if _, err := encoding.Encode("toml", testStruct, encoding.WithIndent("")); err == nil {
		t.Error("toml this option should no supported")
	}
}
