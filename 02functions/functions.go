package functions

import (
	"errors"
	"fmt"
	"strings"
)

func Functions() {

	fmt.Println("-------------------------------- 02 functions starts --------------------------------")

	privateFunction()
	PublicFunction()

	// multiple return values - when got some error, the function will return the error and the result will be 0
	err, num := multipleReturnValues(0, 1)
	fmt.Println("err:", err)
	fmt.Println("num:", num)

	// multiple return values - when got no error, the function will return the error and the result will be the sum of a and b
	err, num = multipleReturnValues(1, 1)
	fmt.Println("err:", err)
	fmt.Println("num:", num)

	fmt.Println("formmatErrorMessages:", formmatErrorMessages("a or b is 0"))

	fmt.Println("-------------------------------- 02 functions ends --------------------------------")

}

func privateFunction() {
	printMessage("privateFunction")
}

func PublicFunction() {
	printMessage("PublicFunction")
}

func printMessage(message string) {
	fmt.Println(message)
}

func multipleReturnValues(a int, b int) (error, int) {
	if a == 0 || b == 0 {
		return errors.New("a or b is 0"), 0
	}
	return nil, a + b
}

func formmatErrorMessages(message string) string {
	m := strings.ReplaceAll(message, " ", "_")
	m = strings.ToUpper(m)
	return m
}
