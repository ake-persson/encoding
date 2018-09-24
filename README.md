[![GoDoc](https://godoc.org/github.com/mickep76/encoding?status.svg)](https://godoc.org/github.com/mickep76/encoding)
[![codecov](https://codecov.io/gh/mickep76/encoding/branch/master/graph/badge.svg)](https://codecov.io/gh/mickep76/encoding)
[![Build Status](https://travis-ci.org/mickep76/encoding.svg?branch=master)](https://travis-ci.org/mickep76/encoding)
[![Go Report Card](https://goreportcard.com/badge/github.com/mickep76/encoding)](https://goreportcard.com/report/github.com/mickep76/encoding)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mickep76/mlfmt/blob/master/LICENSE)

# encoding

Package provides a generic interface to encoders and decoders

## Example

More examples can be found in [examples](https://github.com/mickep76/encoding/tree/master/examples).

```go
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

        b, err := encoding.ToBytes("yaml", in)
        if err != nil {
                log.Fatal(err)
        }

        out := Messages{}
        if err := encoding.FromBytes("yaml", b, &out); err != nil {
                log.Fatal(err)
        }

        for _, m := range out {
                fmt.Printf("%s: %s\n", m.Name, m.Text)
        }
}
```
