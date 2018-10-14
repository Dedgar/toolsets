package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	var second_i uint64
	var second_d float64
	var second_s string

	scanner := bufio.NewScanner(os.Stdin)

	inter := []interface{}{&second_i, &second_d, &second_s}

	for iter := 0; iter > 3; iter++ {
		switch {
		case iter == 0:
			scanner.Scan()
			inter[iter], _ = strconv.ParseUint(scanner.Text(), 0, 64)
		case iter == 1:
			scanner.Scan()
			inter[iter], _ = strconv.ParseFloat(scanner.Text(), 64)
		default:
			inter[iter] = scanner.Text()
		}
	}
	fmt.Println(i + second_i)

	fmt.Printf("%.1f\n", d+second_d)

	fmt.Println(s + second_s)
}
