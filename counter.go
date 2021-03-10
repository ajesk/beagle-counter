package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/ajesk/beagle-counter/gpio"
)

var bits = map[int]bool{
	64: false,
	32: false,
	16: false,
	8:  false,
	4:  false,
	2:  false,
	1:  false,
}
var counter int = 0

func main() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("\n%d\t", counter)
		decompose(counter)
		disperseValues()
		counter++
	}
}

func decompose(num int) {
	var calc = num
	for calc != 0 {
		for _, v := range sortedKeys(bits) {
			if calc >= v {
				bits[v] = true
				calc = calc - v
			} else {
				bits[v] = false
			}
		}
	}
}

func sortedKeys(m map[int]bool) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func disperseValues() {
	for k, v := range bits {
		if v {
			fmt.Printf("%d ", k)
		}
	}
}

func setGpio() {
	pins := gpio.GPIO{}

	pins.Pin("2")
}
