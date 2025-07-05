package types

import (
	"fmt"
	"unicode/utf8"
)

func Types() {
	var a int
	var b int32
	var c int64

	a = 10
	b = 20
	c = 30

	fmt.Println("-------------------------------- 01 types --------------------------------")

	fmt.Println("a:", a, "b:", b, "c:", c)

	const d, e = 10, 20

	fmt.Println("d:", d, "e:", e)

	var myString string = "resumé" // its not 6 bytes, its 7 bytes

	fmt.Println("Length of myString:", len(myString))
	fmt.Println("Rune count of myString:", utf8.RuneCountInString(myString))

	var myBool bool = true

	fmt.Println("myBool:", myBool)

	var myFloat float32 = 1.234567890

	fmt.Println("myFloat:", myFloat)

	var myComplex complex64 = 1.234567890 + 1.234567890i

	fmt.Println("myComplex:", myComplex)

	var myRune rune = 'a'

	fmt.Println("myRune:", myRune)

	fmt.Println("-------------------------------- 01 types ends --------------------------------")

}
