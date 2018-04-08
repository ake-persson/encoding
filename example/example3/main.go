package main

import (
	"bufio"
	"log"
	"os"

	"github.com/mickep76/encdec"
	_ "github.com/mickep76/encdec/bson"
)

type Message struct {
	Name, Text string
}

type Messages []*Message

func main() {
	msgs := Messages{
		&Message{Name: "Ed", Text: "Knock knock."},
		&Message{Name: "Sam", Text: "Who's there?"},
		&Message{Name: "Ed", Text: "Go fmt."},
		&Message{Name: "Sam", Text: "Go fmt who?"},
		&Message{Name: "Ed", Text: "Go fmt yourself!"},
	}

	fp, err := os.Create("example3.bson")
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(fp)
	enc, err := encdec.NewEncoder("bson", w)
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range msgs {
		enc.Encode(msg)
	}
	w.Flush()
}
