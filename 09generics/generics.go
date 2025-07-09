package generics

import "fmt"

func Generics() {
	fmt.Println("-------------------------------- 09 generics starts --------------------------------")

	var numbers = []int{1, 2, 3, 4, 5}
	var floats = []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(squareSlice(numbers))
	fmt.Println(squareSlice(floats))

	fmt.Println("-------------------------------- 09 generics ends --------------------------------")
}

func squareSlice[T int | float64](slice []T) []T {
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}
