package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/mickep76/encoding"
	_ "github.com/mickep76/encoding/json"
	_ "github.com/mickep76/encoding/toml"
	_ "github.com/mickep76/encoding/yaml"
)

type Message struct {
	Name, Text string
}

type Messages struct {
	Messages []*Message
}

func main() {
	format := flag.String("fmt", "json", fmt.Sprintf("Encodings: [%s].", strings.Join(encoding.Encodings(), ", ")))
	indent := flag.String("indent", "", "Indent encoding (only supported by JSON)")

	flag.Parse()

	in := Messages{
		Messages: []*Message{
			&Message{Name: "Ed", Text: "Knock knock."},
			&Message{Name: "Sam", Text: "Who's there?"},
			&Message{Name: "Ed", Text: "Go fmt."},
			&Message{Name: "Sam", Text: "Go fmt who?"},
			&Message{Name: "Ed", Text: "Go fmt yourself!"},
		},
	}

	var opts []encoding.Option
	if *indent != "" {
		opts = append(opts, encoding.WithIndent(*indent))
	}
	e, err := encoding.NewEncoding(*format, opts...)

	b, err := e.Encode(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Encoded:\n%s\n", string(b))

	out := Messages{}
	if err := e.Decode(b, &out); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decoded:")
	for _, m := range out.Messages {
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
