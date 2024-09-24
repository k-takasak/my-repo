package main

import "fmt"

var version string

func main() {
	fmt.Printf("Example %s\n", version)
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}