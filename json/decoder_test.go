package json

import (
	"testing"

	"github.com/go-test/deep"

	"github.com/mickep76/encoding"
)

func TestDecode(t *testing.T) {
	var g interface{}
	if err := encoding.Decode("json", []byte(testEncoded), &g); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMap); err != nil {
		t.Error(err)
	}
}

func TestDecodeWithMapString(t *testing.T) {
	var g interface{}
	if err := encoding.Decode("json", []byte(testEncoded), &g, encoding.WithMapString()); err == nil {
		t.Error("this option should not be supported for json")
	}
}
