package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	testCase := make(map[int]int)

	testCase[0] = 0
	testCase[5] = 5
	testCase[10] = 55
	testCase[20] = 6765

	for key, value := range testCase{
		if got := Fib(key); got != value{
			fmt.Printf("Failed send: %d, got: %d, expected: %d", key, got, value)
			t.Fail()
		}
	}

	fmt.Println("Success!")
}
