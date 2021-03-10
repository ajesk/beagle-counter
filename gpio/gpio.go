package gpio

import (
	"io/ioutil"
	"os"
)

// GPIO placeholder
type GPIO struct{}

// Pin generate a pin
func (r GPIO) Pin(name string) GPIOPin {
	pin := GPIOPin{name}
	filename := pin.Filename()
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		ioutil.WriteFile("/sys/class/gpio/export", []byte(pin.Name), 0666)
	}
	return pin
}

// GPIOPin pin name
type GPIOPin struct {
	Name string
}

// Filename specifies path to GPIO
func (r GPIOPin) Filename() string {
	return "/sys/class/gpio/gpio" + r.Name
}

func (r GPIOPin) write(where, what string) GPIOPin {
	filename := r.Filename() + "/" + where
	ioutil.WriteFile(filename, []byte(what), 0666)
	return r
}

// Output sets a given pin to out
func (r GPIOPin) Output() GPIOPin {
	return r.write("direction", "out")
}

// Input sets a given pin to in
func (r GPIOPin) Input() GPIOPin {
	return r.write("direction", "in")
}

// High sets pin to high
func (r GPIOPin) High() GPIOPin {
	return r.write("value", "1")
}

// Low sets pin to low
func (r GPIOPin) Low() GPIOPin {
	return r.write("value", "0")
}
