package arraysslicesmapsloops

import "fmt"

func ArraysSlicesMapsLoops() {
	fmt.Println("-------------------------------- 03 arrays_slices_maps_loops starts --------------------------------")

	arraysAndSlices()
	maps()
	loops()

	fmt.Println("-------------------------------- 03 arrays_slices_maps_loops ends --------------------------------")
}
func arraysAndSlices() {
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

	// unknown size array
	unknownSizeArray := [...]int{1, 2, 3, 4, 5}
	fmt.Println("unknownSizeArray:", unknownSizeArray)

	// multi-dimensional array
	multiDimensionalArray := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("multiDimensionalArray:", multiDimensionalArray)
}

func maps() {
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	fmt.Println("map1:", map1)

	map2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("map2:", map2)

	// we can get the value of a key and if the key exists
	// this example exists will be true
	var valueOfA, exists = map1["a"]
	fmt.Println("valueOfA:", valueOfA, "exists:", exists)

	// this example exists will be false
	var valueOfD, ok = map1["d"]
	fmt.Println("valueOfD:", valueOfD, "exists:", ok)

	// delete a key
	delete(map1, "a")
	fmt.Println("map1:", map1)

	// iterate over a map
	for key, value := range map1 {
		fmt.Println("key:", key, "value:", value)
	}
}

func loops() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	// while loop
	index := 0
	for index < 10 {
		fmt.Println("index:", index)
		index++
	}

	for {
		if index >= 20 {
			break
		}
		fmt.Println("index:", index)
		index++
	}

	// continue
	for index := 0; index < 10; index++ {
		if index == 5 {
			continue // skip the current iteration
		}
		fmt.Println("index:", index)
	}

}
