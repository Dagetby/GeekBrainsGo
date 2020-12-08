package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {

	a := generateArray(100)

	resultInsert := SortInsert(a)
	resultBubble := SortBubble(a)
	sort.Ints(a)

	for i, v := range a {
		if v != resultBubble[i] {
			fmt.Printf("Fail bubble sort: index- %d, expected array: %d, got: %d ", i, a, resultBubble)
			t.Fail()
		}
		if v != resultInsert[i] {
			fmt.Printf("Fail Insert sort: index- %d, expected array: %d, got: %d ", i, a, resultBubble)
			t.Fail()
		}
	}

	fmt.Println("Success!")
}

func generateArray(lenArray int) []int {

	var a = make([]int, 0, lenArray)

	for i := 0; i < lenArray; i++ {
		a = append(a, rand.Int())
	}
	return a
}

func BenchmarkInsert(b *testing.B) {

	a := generateArray(100000)
	SortInsert(a)
}

func BenchmarkBubble(b *testing.B) {

	a := generateArray(100000)
	SortBubble(a)
}
