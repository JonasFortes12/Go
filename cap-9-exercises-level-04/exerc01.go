package cap9exerciseslevel04

import "fmt"

func Exerc01() {

	/* var arr [5]int = [5]int{1, 2, 3, 4, 5} */

	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for _, value := range arr {
		fmt.Println(value)
	}

	fmt.Printf("Array: %v | Type: %T", arr, arr)

}
