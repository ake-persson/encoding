package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mickep76/encoding"
	_ "github.com/mickep76/encoding/toml"
)

type Message struct {
	Name, Text string
}

type Messages struct {
	Messages []*Message
}

func EncToFile(fn string, name string, v interface{}) error {
	fp, err := os.Create(fn)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(fp)
	enc, err := encoding.NewEncoder(name, w)
	if err != nil {
		return err
	}

	if err := enc.Encode(v); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	if err := fp.Close(); err != nil {
		return err
	}

	return nil
}

func DecFromFile(fn string, name string, v interface{}) error {
	fp, err := os.Open(fn)
	if err != nil {
		return err
	}

	r := bufio.NewReader(fp)
	dec, err := encoding.NewDecoder(name, r)
	if err != nil {
		return err
	}

	if err := dec.Decode(v); err != nil {
		return err
	}

	if err := fp.Close(); err != nil {
		return err
	}

	return nil
}

func main() {
	in := Messages{
		Messages: []*Message{
			&Message{Name: "Ed", Text: "Knock knock."},
			&Message{Name: "Sam", Text: "Who's there?"},
			&Message{Name: "Ed", Text: "Go fmt."},
			&Message{Name: "Sam", Text: "Go fmt who?"},
			&Message{Name: "Ed", Text: "Go fmt yourself!"},
		},
	}

	if err := EncToFile("example3.toml", "toml", in); err != nil {
		log.Fatal(err)
	}

	out := Messages{}
	if err := DecFromFile("example3.toml", "toml", &out); err != nil {
		log.Fatal(err)
	}

	for _, m := range out.Messages {
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
