package main

import (
	"testing"
)

func expressionsSum(args ...int) func(...int) int {
	sum := 1
	for _, v := range args {
		sum = sum * v
	}

	if len(args) == 0 {
		return func(args ...int) int {
			return sum
		}
	}

	return func(args ...int) int {
		for _, v := range args {
			sum = sum * v
		}
		return sum
	}
}

func TestExpressionsSum(t *testing.T) {
	if expressionsSum(2, 3)() != 6 {
		t.Fatal()
	}
	if expressionsSum(2)(3) != 6 {
		t.Fatal()
	}
}
