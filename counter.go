package main

import (
	"fmt"
	"time"

	"github.com/ajesk/beagle-counter/gpio"
)

var defaultBits = map[int]bool{
	64: false,
	32: false,
	16: false,
	8:  false,
	4:  false,
	2:  false,
	1:  false,
}

var pins = map[int]int{
	64: 66,
	32: 65,
	16: 46,
	8:  26,
	4:  44,
	2:  68,
	1:  67,
}

var counter int = 0

func main() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("\n%d\t", counter)
		binary := toBinary(counter)
		bits := flipBits(binary)
		fmt.Print(bits)
		disperseValues()
		counter++
	}
}

func increment() {
	counter++
	if counter > 255 {
		counter = 0
	}
}

func toBinary(num int) []bool {
	var binary []bool
	var calc = num
	for calc != 0 {
		binary = append(binary, calc%2 != 0)
		calc = calc / 2
	}

	return binary
}

func flipBits(binary []bool) map[int]bool {
	bitSet := copyBits()
	count := 1
	for _, bit := range binary {
		bitSet[count] = bit
		count *= 2
	}
	return bitSet
}

func copyBits() (copy map[int]bool) {
	copy = map[int]bool{}
	for k, v := range defaultBits {
		copy[k] = v
	}
	return copy
}

func disperseValues() {
	for k, v := range defaultBits {
		if v {
			fmt.Printf("%d ", k)
		}
	}
}

func setGpio() {
	pins := gpio.GPIO{}

	pins.Pin("2")
}
