package main

import (
	"fmt"
)

func main() {
	num := 75
	switch {
	case num%3 == 0 && num%5 == 0:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

}
