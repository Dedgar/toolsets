package main

import "fmt"
import "strings"

func main() {
	if strings.Contains("original", "orig") && strings.Contains("original", "not") == false {
		fmt.Println("yep!")

	} else {
		fmt.Println("nope!")
	}
}
