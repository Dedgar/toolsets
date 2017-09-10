package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	samp_str := "%e4%b9%9d"
	fmt.Println(samp_str)
	test_str, err := url.QueryUnescape(samp_str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(test_str)

	new_str, err := url.QueryUnescape("九")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(new_str)
}
