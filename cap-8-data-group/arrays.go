package cap8datagroup

import "fmt"

type arrayOfTen [10]int

func Arrays() {

	// Arrays are fixed size, so we need to define the size of the array when we declare it
	var x [5]int
	var y arrayOfTen

	fmt.Printf("X: %v | Type %T\n Y: %v | Type %T", x, x, y, y)

	y[9] = 1

	array := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("\nArray: %v | Type %T", array, array)

	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("\nSlice: %v | Type %T", slice, slice)

	slice = append(slice, 6)

	fmt.Printf("\nSlice: %v | Type %T", slice, slice)

	//|This op eration will cause a panic because the index is out of range
	//slice[20] = 20

	//|This operation will not cause a panic because the slice will be resized
	slice = append(slice, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	slice[20] = 20
}
