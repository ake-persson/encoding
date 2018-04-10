package json

import (
	"testing"

	"github.com/mickep76/encdec"
)

var (
	testJSON = `{"test":{"abc":{"myTest":"Test2"},"myTest":"Test"}}
`

	testJSONIndent = `{
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

func TestToByte(t *testing.T) {
	b, err := encdec.ToBytes("json", &testMap)
	if err != nil {
		t.Error(err)
	}

	if string(b) != testJSON {
		t.Errorf("want:\n%s, got:\n%s", testJSON, string(b))
	}
}
