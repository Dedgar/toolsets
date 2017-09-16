package main

import (
	"fmt"
)

func main() {
	selection := "grade"
	level := 1
	entry := map[string]string{"kanji": "一", "von": "イチ,イツ,"}
	entrymap := map[string]interface{}{"entry": entry, "selection": selection, "level": level}

	fmt.Println(entrymap)
	fmt.Println(entrymap["entry"])
	fmt.Println(entrymap["selection"])
}
