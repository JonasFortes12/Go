package cap9exerciseslevel04

import "fmt"

func Exerc02() {

	//slice := []int{}
	//var slice []int = make([]int, 5)
	//var slice []int = []int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("Array: %v | Type: %T", slice, slice)

}
