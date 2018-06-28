package main

import (
	"fmt"
)

func fizzbuzz(number int) {
	switch {
	case number%3 == 0 && number%5 == 0:
		fmt.Println(number, "FizzBuzz")
	case number%3 == 0:
		fmt.Println(number, "Fizz")
	case number%5 == 0:
		fmt.Println(number, "Buzz")
	}
}

func main() {
	start_num := 1

	for start_num < 21 {
		fizzbuzz(start_num)
		start_num++
	}
}
