package main

import (
	"fmt"
	"log"

	"github.com/mickep76/encoding"
	_ "github.com/mickep76/encoding/yaml"
)

type Message struct {
	Name, Text string
}

type Messages []*Message

func main() {
	in := Messages{
		&Message{Name: "Ed", Text: "Knock knock."},
		&Message{Name: "Sam", Text: "Who's there?"},
		&Message{Name: "Ed", Text: "Go fmt."},
		&Message{Name: "Sam", Text: "Go fmt who?"},
		&Message{Name: "Ed", Text: "Go fmt yourself!"},
	}

	b, err := encoding.Encode("yaml", in)
	if err != nil {
		log.Fatal(err)
	}

	out := Messages{}
	if err := encoding.Decode("yaml", b, &out); err != nil {
		log.Fatal(err)
	}

	for _, m := range out {
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
