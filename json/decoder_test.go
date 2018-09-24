package json

import (
	"testing"

	"github.com/go-test/deep"

	"github.com/mickep76/encoding"
)

func TestFromByte(t *testing.T) {
	var g interface{}
	if err := encoding.FromBytes("json", []byte(testJSON), &g); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMap); err != nil {
		t.Error(err)
	}
}

func TestFromByteWithMapString(t *testing.T) {
	var g interface{}
	if err := encoding.FromBytes("json", []byte(testJSON), &g, encoding.WithMapString()); err != nil {
		t.Error(err)
	}
}

func TestFromFile(t *testing.T) {
	var g interface{}
	if err := encoding.FromFile("json", "test.json", &g); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMap); err != nil {
		t.Error(err)
	}
}
