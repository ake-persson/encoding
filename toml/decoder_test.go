package toml

import (
	"testing"

	"github.com/go-test/deep"

	"github.com/mickep76/encoding"
)

func TestDecode(t *testing.T) {
	g := Messages{}
	if err := encoding.Decode("toml", []byte(testEncoded), &g); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testStruct); err != nil {
		t.Error(err)
	}
}

func TestDecodeWithMapString(t *testing.T) {
	g := Messages{}
	if err := encoding.Decode("toml", []byte(testEncoded), &g, encoding.WithMapString()); err != nil {
		t.Error(err)
	}
}
