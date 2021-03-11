package main

import (
	"reflect"
	"testing"
)

func Test122(t *testing.T) {
	expected := []bool{false, true, false, true, true, true, true}
	result := decompose(122)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test0(t *testing.T) {
	expected := []bool{}
	result := decompose(0)

	if len(expected) != len(result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}

func Test1(t *testing.T) {
	expected := []bool{true}
	result := decompose(1)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\nexpected: %v\nreceived: %v", expected, result)
	}
}
