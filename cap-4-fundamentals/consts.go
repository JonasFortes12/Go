package cap4fundamentals

import "fmt"

const numConst = 10

var numVar = 10

var num int
var num2 float64

func Constants() {

	// This has a error because numVar already has the defined type as int
	// num2 = numVar

	// This works because numConst does not have a defined type while it is not used
	num2 = numConst

	fmt.Printf("Value: %v\t Type: %T", num2, num2)

}
