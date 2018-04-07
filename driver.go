package encdec

var drivers = make(map[string]Driver)

// Driver interface.
type Driver interface {
	SetIndent(indent string) error
	Encoder() (Encoder, error)
	Decoder() (Decoder, error)
}

// Register driver.
func Register(name string, driver Driver) {
	drivers[name] = driver
}
