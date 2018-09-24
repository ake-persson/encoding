package toml

import (
	"io/ioutil"
	"os"
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
	testTOML = `[[Messages]]
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

	tmpFile = "test.tmp"
)

func readFile(fn string) ([]byte, error) {
	fp, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	if err := fp.Close(); err != nil {
		return nil, err
	}

	return b, nil
}

func TestToByte(t *testing.T) {
	b, err := encoding.ToBytes("toml", testStruct)
	if err != nil {
		t.Error(err)
	}

	if string(b) != testTOML {
		t.Errorf("want:\n%s\n\ngot:\n%s", testTOML, string(b))
	}
}

func TestToByteWithIndent(t *testing.T) {
	if _, err := encoding.ToBytes("toml", testStruct, encoding.WithIndent("")); err != nil {
		t.Error(err)
	}
}

func TestToFile(t *testing.T) {
	if err := encoding.ToFile("toml", tmpFile, testStruct); err != nil {
		t.Error(err)
	}

	b, err := readFile(tmpFile)
	if err != nil {
		t.Error(err)
	}

	if string(b) != testTOML {
		t.Errorf("want:\n%s\n\ngot:\n%s", testTOML, string(b))
	}

	if err := os.Remove(tmpFile); err != nil {
		t.Error(err)
	}
}
