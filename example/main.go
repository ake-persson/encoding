package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/mickep76/encdec"
	_ "github.com/mickep76/encdec/json"
)

func main() {
	const data = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}

	dec, err := encdec.NewDecoder("json", strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
