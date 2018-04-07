package encdec

import "fmt"

// Decoder interface.
type Decoder interface {
	Decode(bytes []byte, value interface{}) error
}

// Decoder constructor.
func NewDecoder(driver string, options ...func(Driver) error) (Decoder, error) {
	d, ok := drivers[driver]
	if !ok {
		return nil, fmt.Errorf("driver is not registered: %s", driver)
	}

	for _, option := range options {
		if err := option(d); err != nil {
			return nil, err
		}
	}

	return d.Decoder(), nil
}
