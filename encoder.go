package encdec

import (
	"fmt"
	"io"
)

// Encoder interface.
type Encoder interface {
	Encode(value interface{}) error
}

// NewEncoder constructor.
func NewEncoder(driver string, writer io.Writer, options ...func(Driver) error) (Encoder, error) {
	d, ok := drivers[driver]
	if !ok {
		return nil, fmt.Errorf("driver is not registered: %s", driver)
	}

	for _, option := range options {
		if err := option(d); err != nil {
			return nil, err
		}
	}

	return d.Encoder()
}

// Indent output.
func Indent(indent string) func(Driver) error {
	return func(d Driver) error {
		return d.SetIndent(indent)
	}
}
