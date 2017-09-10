package main

import (
	"fmt"
)

//type Kanji struct {
//	kanj   string
//	von    string
//	vkun   string
//	transl string
//	roma   string
//	rememb string
//	jlpt   string
//	school string
//}

//type Layer struct {
//	midlayer map[string]string
//}

func main() {
	entry := make(map[string]map[string]string)
	//  newent := make(map[string]map[string]string)

	//	newent["一"] = {"一", "イチ,イツ,", "ひと-,ひと.つ", "one", "ichi,itsu hito-,hito.tsu", "1", "5", "1"}
	entry["一"] = map[string]string{"kanji": "一", "von": "イチ,イツ,", "vkun": "ひと-,ひと.つ", "transl": "one", "roma": "ichi,itsu hito-,hito.tsu", "rememb": "1", "jlpt": "5", "school": "1"}
	entry["二"] = map[string]string{"kanji": "一", "von": "イチ,イツ,", "vkun": "ひと-,ひと.つ", "transl": "one", "roma": "ichi,itsu hito-,hito.tsu", "rememb": "1", "jlpt": "5", "school": "1"}
	//	entry["二"] = Kanji{"二", "イチ,イツ,", "ひと-,ひと.つ", "one", "ichi,itsu hito-,hito.tsu", "1", "5", "1"}

	fmt.Println(entry)
	fmt.Println(entry["二"])
	fmt.Println(entry["二"]["roma"])

}
