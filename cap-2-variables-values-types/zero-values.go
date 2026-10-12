package cap2variablesvaluestypes

import "fmt"

var a int
var b string
var c bool
var d float64

func ZeroValues() {

	fmt.Printf("Value %v, Type %T \n", a, a)
	fmt.Printf("Value %v, Type %T \n", b, b)
	fmt.Printf("Value %v, Type %T \n", c, c)
	fmt.Printf("Value %v, Type %T \n", d, d)
}
