package yaml

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"

	"github.com/mickep76/encoding"
)

func TestDecode(t *testing.T) {
	var g interface{}
	if err := encoding.Decode("yaml", []byte(testEncoded), &g); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMapInterface); err != nil {
		t.Error(err)
	}
}

func TestDecodeWithMapString(t *testing.T) {
	var g interface{}
	if err := encoding.Decode("yaml", []byte(testEncoded), &g, encoding.WithMapString()); err != nil {
		t.Error(err)
	}

	if err := deep.Equal(g, testMap); err != nil {
		t.Error(err)
	}
}

func TestDecodeWithMapStringError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	g := ""
	if err := encoding.Decode("yaml", []byte(testEncoded), g, encoding.WithMapString()); err != nil {
		t.Error(fmt.Errorf("expected a panic! %v", err))
	}
}
