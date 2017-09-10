package main

import (
	"fmt"
)

type Kanji struct {
	kanj   string
	von    string
	vkun   string
	transl string
	roma   string
	rememb string
	jlpt   string
	school string
}

//type Key struct {
//	Path, Country string
//}

func main() {
	//	hits := make(map[Key]int)

	//	hits[Key{"/", "vn"}]++
	//	hits[Key{"/", "vn"}]++
	//	hits[Key{"/", "us"}]++

	entry := make(map[string]Kanji)

	entry["一"] = Kanji{"一", "イチ,イツ,", "ひと-,ひと.つ", "one", "ichi,itsu hito-,hito.tsu", "1", "5", "1"}
	entry["二"] = Kanji{"二", "イチ,イツ,", "ひと-,ひと.つ", "one", "ichi,itsu hito-,hito.tsu", "1", "5", "1"}

	fmt.Println(entry)
	fmt.Println(entry["二"])

	var slicie []Kanji
	slicie = append(slicie, Kanji{"一", "イチ,イツ,", "ひと-,ひと.つ", "one", "ichi,itsu hito-,hito.tsu", "1", "5", "1"})
	slicie = append(slicie, Kanji{"二", "イチ,イツ,", "ひと-,ひと.つ", "one", "ichi,itsu hito-,hito.tsu", "1", "5", "1"})
	fmt.Println(slicie)
}
