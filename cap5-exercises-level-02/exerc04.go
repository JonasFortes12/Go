package cap5exerciseslevel02

import "fmt"

func Exerc04() {

	num := 122

	fmt.Printf("Value: %v | Decimal: %d \n", num, num)
	fmt.Printf("Value: %v | Binary: %b \n", num, num)
	fmt.Printf("Value: %v | Hexadecimal: %#x \n", num, num)

	num2 := num << 1

	fmt.Printf("Value: %v | Decimal: %d \n", num2, num2)
	fmt.Printf("Value: %v | Binary: %b \n", num2, num2)
	fmt.Printf("Value: %v | Hexadecimal: %#x \n", num2, num2)
}
