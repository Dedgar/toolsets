package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'

	str1, str2 := strings.Split(str, "\n")[0], strings.Split(str, "\n")[1]

	fmt.Println(str1)
	fmt.Println(str2)
}
