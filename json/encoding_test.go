package json

import (
	"fmt"
	"testing"

	"github.com/mickep76/encoding"
)

func inList(a string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}

func TestEncodings(t *testing.T) {
	encs := encoding.Encodings()

	if !inList("json", encs) {
		t.Error(fmt.Errorf("encoding not in list"))
	}

	if ok := encoding.Registered("json"); !ok {
		t.Error("not registered")
	}
}
