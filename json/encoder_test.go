package json

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mickep76/encoding"
)

func TestEncode(t *testing.T) {
	b, err := encoding.ToBytes("json", testMap)
	if err != nil {
		t.Error(err)
	}

	if string(b) != testJSON {
		t.Errorf("want:\n%s, got:\n%s", testJSON, string(b))
	}
}

func TestEncodeWithIndent(t *testing.T) {
	b, err := encoding.ToBytes("json", testMap, encoding.WithIndent("  "))
	if err != nil {
		t.Error(err)
	}

	if string(b) != testJSONIndent {
		t.Errorf("want:\n%s, got:\n%s", testJSONIndent, string(b))
	}
}
