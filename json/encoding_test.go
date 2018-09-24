package json

import (
	"fmt"
	"testing"

	"github.com/mickep76/encoding"
)

var (
	testEncoded = `{"test":{"abc":{"myTest":"Test2"},"myTest":"Test"}}
`

	testEncodedIndent = `{
  "test": {
    "abc": {
      "myTest": "Test2"
    },
    "myTest": "Test"
  }
}
`

	testMap = map[string]interface{}{
		"test": map[string]interface{}{
			"myTest": "Test",
			"abc": map[string]interface{}{
				"myTest": "Test2",
			},
		},
	}
)

func inList(a string, l []string) bool {
	for _, b := range l {
		if a == b {
			return true
		}
	}
	return false
}

func TestEncodings(t *testing.T) {
	if !inList("json", encoding.Encodings()) {
		t.Error(fmt.Errorf("json encoding not in registered encodings"))
	}

	if _, err := encoding.Registered("json"); err != nil {
		t.Error(err)
	}
}
