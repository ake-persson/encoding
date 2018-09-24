package yaml

import (
	"fmt"
	"testing"

	"github.com/mickep76/encoding"
)

var (
	testEncoded = `test:
  abc:
    myTest: Test2
  myInt: 1
  myList:
  - Item 1
  - Item 2
  myTest: Test
`

	testMap = map[string]interface{}{
		"test": map[string]interface{}{
			"myTest": "Test",
			"abc": map[string]interface{}{
				"myTest": "Test2",
			},
			"myList": []interface{}{
				"Item 1",
				"Item 2",
			},
			"myInt": 1,
		},
	}

	testMapInterface = map[interface{}]interface{}{
		"test": map[interface{}]interface{}{
			"myTest": "Test",
			"abc": map[interface{}]interface{}{
				"myTest": "Test2",
			},
			"myList": []interface{}{
				"Item 1",
				"Item 2",
			},
			"myInt": 1,
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
	if !inList("yaml", encoding.Encodings()) {
		t.Error(fmt.Errorf("yaml encoding not in registered encodings"))
	}

	if _, err := encoding.Registered("yaml"); err != nil {
		t.Error(err)
	}
}
