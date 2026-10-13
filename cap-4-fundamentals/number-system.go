package cap4fundamentals

import "fmt"

func NumberSystem() {

	number := 122

	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Hexadecimal: %x\n", number)

	//122 - 64 == 58 - 32 == 26 - 16 == 10 - 8 == 2

	//122 - 112 == 10

	// For num 122
	// ... 2^6 2^5 2^4 2^3 2^2 2^1 2^0
	//      1   1   1   1   0	1   0

	// ... 16^4 16^3 16^2 16^1 16^0
	//                      7   A
}
