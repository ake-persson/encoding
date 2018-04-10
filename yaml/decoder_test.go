package yaml

import (
	"testing"

	"github.com/go-test/deep"

	"github.com/mickep76/encdec"
)

func TestFromByte(t *testing.T) {
	var g interface{}
	if err := encdec.FromBytes("yaml", []byte(testYAML), &g); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMapInterface); err != nil {
		t.Error(err)
	}
}

func TestFromByteWithMapString(t *testing.T) {
	var g interface{}
	if err := encdec.FromBytes("yaml", []byte(testYAML), &g, encdec.WithMapString()); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMap); err != nil {
		t.Error(err)
	}
}

func TestFromFile(t *testing.T) {
	var g interface{}
	if err := encdec.FromFile("yaml", "test.yaml", &g); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMapInterface); err != nil {
		t.Error(err)
	}
}
