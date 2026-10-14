package cap6ControlFlow

import "fmt"

func Switch() {

	num := 'b'

	switch num {
	case 1:
		println("num is equal to 1")
	case 2:
		println("num is equal to 2")
	case 3:
		println("num is equal to 3")
	case 4:
		println("num is equal to 4")
		fallthrough // Fall through to the next case without checking the condition
	case 'a':
		println("num is equal to 'a' or fallthrough to this case")

	case 'b', 'c':
		println("num is equal to 'b' or 'c'")
	default:
		println("num is not equal to 1, 2, 3 or 4, it's equal to", num)
	}

	x := -100

	switch {
	case x < 5:
		println("x is less than 5")
	case x >= 5 && x <= 10:
		println("x is between 5 and 10")
	case x >= 5, x <= 10:
		println("x is greater than or equal to 5 or x is less than or equal to 10")
	default:
		println("x is greater than 10")
	}

	var data interface{}
	data = -10

	switch data.(type) {
	case int:
		fmt.Println("Integer")
	case bool:
		fmt.Println("Boolean")
	case float64:
		fmt.Println("Float64")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Type not found")
	}

	switch x := 10; {
	case x == 10:
		fmt.Println("X is 10")
	default:
		fmt.Println("Not is 10")

	}

}
