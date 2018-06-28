package main

import (
	"fmt"
	"os"
)

func availableVids(show string, season string, episode string) bool {
	teststr := "./:static/vid/" + show + "/" + season + "/" + episode + ".mp4"
	fmt.Println(teststr)
	if _, err := os.Stat(teststr); err == nil {
		return true
	}
	return false
}

func main() {
	mybool := availableVids("divisionrune", "1", "1")
	fmt.Println(mybool)
}
