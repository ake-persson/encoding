package toml

import (
	"fmt"
	"testing"

	"github.com/mickep76/encoding"
)

type Message struct {
	Name, Text string
}

type Messages struct {
	Messages []*Message
}

var (
	testEncoded = `[[Messages]]
  Name = "Ed"
  Text = "Knock knock."

[[Messages]]
  Name = "Sam"
  Text = "Who's there?"

[[Messages]]
  Name = "Ed"
  Text = "Go fmt."

[[Messages]]
  Name = "Sam"
  Text = "Go fmt who?"

[[Messages]]
  Name = "Ed"
  Text = "Go fmt yourself!"
`

	testStruct = Messages{
		Messages: []*Message{
			&Message{Name: "Ed", Text: "Knock knock."},
			&Message{Name: "Sam", Text: "Who's there?"},
			&Message{Name: "Ed", Text: "Go fmt."},
			&Message{Name: "Sam", Text: "Go fmt who?"},
			&Message{Name: "Ed", Text: "Go fmt yourself!"},
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
	if !inList("toml", encoding.Encodings()) {
		t.Error(fmt.Errorf("toml encoding not in registered encodings"))
	}

	if _, err := encoding.Registered("toml"); err != nil {
		t.Error(err)
	}
}
