package main

import "fmt"

func main() {
	//sort
	arr := []int{4, 1, 8, 9, 10, 213, 41, 12, 22, 3, 45, 6}
	fmt.Println(SortInsert(arr))
	fmt.Println(SortBubble(arr))

	//FizzBuzz
	FizzBuzz()
}

//FizzBuzz...
//Если число кратно 3 выводит Fizz
//Если число кратно 5 выводит Buzz
//А если число кратно и 3 и 5 выводи FizzBuzz
func FizzBuzz() {

	for i := 1; i <= 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")

		case i%3 == 0:
			fmt.Println("Fizz")

		case i%5 == 0:
			fmt.Println("Buzz")

		default:
			fmt.Println(i)
		}
	}
}
