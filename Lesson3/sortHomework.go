package main

func SortInsert(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		m := min(arr[:], i)
		arr[i], arr[m] = arr[m], arr[i]
	}
	return arr
}

func min(arr []int, n int) int {
	for i := n; i < len(arr); i++ {
		if arr[i] < arr[n] {
			n = i
		}
	}
	return n
}

func SortBubble(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
