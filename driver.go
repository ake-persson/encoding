package encdec

var drivers = make(map[string]Driver)

// Driver interface.
type Driver interface {
	Encoder() Encoder
	Decoder() Decoder
}

// Register driver.
func Register(name string, driver Driver) {
	drivers[name] = driver
}
