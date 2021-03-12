package main

import (
	"reflect"
	"testing"
)

func Test_toBinary122(t *testing.T) {
	expected := []bool{false, true, false, true, true, true, true}
	result := toBinary(122)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test_toBinary0(t *testing.T) {
	expected := []bool{}
	result := toBinary(0)

	if len(expected) != len(result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test_toBinary1(t *testing.T) {
	expected := []bool{true}
	result := toBinary(1)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test_flipBits13(t *testing.T) {
	bitArr := []bool{true, true, false, true}
	var expected = map[int]bool{
		64: false,
		32: false,
		16: false,
		8:  true,
		4:  false,
		2:  true,
		1:  true,
	}

	result := flipBits(bitArr)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test_flipBits255(t *testing.T) {
	bitArr := []bool{true, true, true, true, true, true, true}
	var expected = map[int]bool{
		64: true,
		32: true,
		16: true,
		8:  true,
		4:  true,
		2:  true,
		1:  true,
	}

	result := flipBits(bitArr)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test_flipBits0(t *testing.T) {
	bitArr := []bool{}
	var expected = map[int]bool{
		64: false,
		32: false,
		16: false,
		8:  false,
		4:  false,
		2:  false,
		1:  false,
	}

	result := flipBits(bitArr)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test_copyBitsShouldGetBaseBits(t *testing.T) {
	var expected = map[int]bool{
		64: false,
		32: false,
		16: false,
		8:  false,
		4:  false,
		2:  false,
		1:  false,
	}

	result := copyBits()
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test_counterStartsAt0(t *testing.T) {
	if 0 != counter {
		t.Error("not 0")
	}
}

func Test_incrementShouldIncrement(t *testing.T) {
	counter = 0
	increment()
	if 1 != counter {
		t.Error("not 1")
	}
}

func Test_increment256ShouldRollover(t *testing.T) {
	counter = 0
	for i := 0; i < 256; i++ {
		increment()
	}
	if 0 != counter {
		t.Error("did not roll over")
	}
}
