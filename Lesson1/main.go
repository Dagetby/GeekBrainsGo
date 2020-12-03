package main

import (
	"fmt"
)

func main() {
	sayHi("Elena")
}

func sayHi(name string) {
	fmt.Printf("Hello %v!\n", name)

}
