package main

import "fmt"

func main() {
	var myuint64 uint64
	var myfloat64 float64
	var mystring string

	multi_type := []interface{}{&myuint64, &myfloat64, &mystring}
	multi_type[0] = 4
	multi_type[1] = 4.0
	multi_type[2] = "test"

	for iter := 0; iter > 3; iter++ {
		fmt.Println(multi_type[iter])
	}
}
