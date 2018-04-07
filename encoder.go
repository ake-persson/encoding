package encdec

import "fmt"

// Encoder interface.
type Encoder interface {
	Encode(value interface{}) []byte
}

// NewEncoder constructor.
func NewEncoder(driver string, options ...func(Driver) error) (Encoder, error) {
	d, ok := drivers[driver]
	if !ok {
		return nil, fmt.Errorf("driver is not registered: %s", driver)
	}

	for _, option := range options {
		if err := option(d); err != nil {
			return nil, err
		}
	}

	return d.Encoder(), nil
}
