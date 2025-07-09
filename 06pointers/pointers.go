package pointers

import "fmt"

func Pointers() {
	fmt.Println("-------------------------------- 06 pointers starts --------------------------------")

	var p *int = new(int)
	var i int = 42 // 42

	fmt.Println(p)
	fmt.Println(i)

	*p = 10

	fmt.Printf("Value of p: %v	\n", *p)
	fmt.Printf("Value of i: %v\n", i)

	var slice = [3]int{1, 2, 3}
	var newSlice [3]int = square(slice)
	var newSliceUsingPointer [3]int = squareUsingPointer(&slice)

	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(newSliceUsingPointer)

	fmt.Printf("Memory address of slice: %p\n", &slice)
	fmt.Printf("Memory address of newSlice: %p\n", &newSlice)
	fmt.Printf("Memory address of newSliceUsingPointer: %p\n", &newSliceUsingPointer)

	fmt.Println("-------------------------------- 06 pointers ends --------------------------------")
}

func square(slice [3]int) [3]int {
	for i, v := range slice {
		slice[i] = v * v
	}

	return slice
}

func squareUsingPointer(slice *[3]int) [3]int {
	for i, v := range slice {
		slice[i] = v * v
	}

	return *slice
}
