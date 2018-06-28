package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	root := "static/divisionrune"
	mydir := filepath.Dir(root)
	fmt.Println(mydir)
}
