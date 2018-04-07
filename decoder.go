package encdec

import (
	"fmt"
	"io"
)

// Decoder interface.
type Decoder interface {
	Decode(value interface{}) error
}

// Decoder constructor.
func NewDecoder(driver string, reader io.Reader, options ...func(Driver) error) (Decoder, error) {
	d, ok := drivers[driver]
	if !ok {
		return nil, fmt.Errorf("driver is not registered: %s", driver)
	}

	for _, option := range options {
		if err := option(d); err != nil {
			return nil, err
		}
	}

	return d.Decoder()
}
