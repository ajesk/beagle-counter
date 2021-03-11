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

/**
66 - 64
65 - 32
46 - 16
26 - 8
44 - 4
68 - 2
67 - 1
*/
// var pins = map[int]int{
// 	64
// }

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

func decompose(num int) []bool {
	var binary []bool
	var calc = num
	for calc != 0 {
		binary = append(binary, calc%2 != 0)
		calc = calc / 2
	}

	return binary
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
