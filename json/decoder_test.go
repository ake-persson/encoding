package json

import (
	"reflect"
	"testing"

	"github.com/mickep76/encdec"
)

func TestFromByte(t *testing.T) {
	var g interface{}
	if err := encdec.FromBytes("json", []byte(testJSON), &g); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(g, testMap) {
		t.Errorf("want:\n%+v, got:\n%+v", testMap, g)
	}
}
