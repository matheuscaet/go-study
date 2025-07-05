package arraysslicesmapsloops

import "fmt"

func ArraysSlicesMapsLoops() {
	fmt.Println("-------------------------------- 03 arrays_slices_maps_loops starts --------------------------------")

	arrays()

	fmt.Println("-------------------------------- 03 arrays_slices_maps_loops ends --------------------------------")
}

func arrays() {
	var array [5]int32

	fmt.Println("array:", array)

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println("array:", array) // print the array
	fmt.Println(array[1:3])      // print the slice from the array

	slice := array[1:3]                                                             // create a slice from the array
	fmt.Printf("slice: value %v, type %T, capacity %v\n", slice, slice, cap(slice)) // print the value and type of the slice
	// append to slice
	slice = append(slice, 4)
	fmt.Printf("slice size: %v and slice capacity: %v\n", len(slice), cap(slice))

	fmt.Println(&slice[0]) // print the address of the first element of the slice

	floatArray := [3]float64{1.1, 2.2, 3.3}
	fmt.Println("floatArray:", floatArray)

}
