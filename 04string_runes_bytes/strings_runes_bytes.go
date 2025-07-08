package stringrunesbytes

import "fmt"

func StringsRunesBytes() {
	fmt.Println("-------------------------------- 04 string_runes_bytes starts --------------------------------")

	var myString string = "résumé"
	var indexed = myString[0]
	fmt.Printf("indexed: %v, %T\n", indexed, indexed)

	for index, runeValue := range myString {
		fmt.Println(runeValue, index)
	}

	var myRune = []rune("résumé")
	fmt.Println(myRune)

	for i, v := range myRune {
		fmt.Println(i, v)
	}

	fmt.Println("-------------------------------- 04 string_runes_bytes ends --------------------------------")
}
