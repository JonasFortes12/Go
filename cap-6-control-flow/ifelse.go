package cap6ControlFlow

import "fmt"

func IfElse() {

	num := 10

	fmt.Printf("num is equal to %v\n", num)

	if num == 1 {
		fmt.Printf("1 is equal to %v\n", num)
	} else {
		fmt.Printf("1 is not equal to 1, it's equal to %v\n", num)
	}

	if num == 2 {
		fmt.Printf("1 is equal to %v\n", num)
	} else if num == 3 {
		fmt.Printf("1 is equal to %v\n", num)
	} else {
		fmt.Printf("1 is not equal to 2 or 3, it's equal to %v\n", num)
	}

}
