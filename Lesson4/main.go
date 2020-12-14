package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {

	var err error
	var n int

	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan(){
		val := reader.Text()

		n, err = strconv.Atoi(val)
		if err != nil{
			fmt.Println("Invalid value please enter, integer:")
			continue
		}

		break
	}

	fmt.Printf("Fibonacci number: %d\n",Fib(n))
}

func Fib(n int) int {
	if n == 1  { return 1 }
	if n == 0  { return 0 }
	return Fib(n-2) + Fib(n-1)
}
