package cap4fundamentals

import (
	"fmt"
	"runtime"
)

func RuneAndBytes() {
	a := 'e'
	b := 'é'
	c := '€'

	// This chars is called rune, and is a alias for int32, and is used to represent a Unicode code point.
	fmt.Printf("%v, %v, %v\n", a, b, c)
	fmt.Printf("%v, %v, %v\n", string(a), string(b), string(c))

	// I will convert the rune to byte, and will print the value of the byte, and the string representation of the byte.
	aByte := []byte(string(a))
	bByte := []byte(string(b))
	cByte := []byte(string(c))

	fmt.Printf("%v, %v\n", aByte, string(aByte)) // 1 Byte to represent the char 'e'
	fmt.Printf("%v, %v\n", bByte, string(bByte)) // 2 Bytes to represent the char 'é'
	fmt.Printf("%v, %v\n", cByte, string(cByte)) // 3 Bytes to represent the char '€'

	//The default type from int and float depends on the architecture of the machine.

	machineArch := runtime.GOARCH
	system := runtime.GOOS
	fmt.Printf("Machine architecture: %v\n", machineArch)
	fmt.Printf("System: %v\n", system)

	// If we have a int and we atribute a larger value than the max value of the int,
	// we will have an overflow and the value will be negative.
	// it will 'rotate' the value to the negative side of the int range.
	var intValue int32 = 2147483647 // Max value for int32
	fmt.Printf("Int value: %v\n", intValue)
	intValue++
	fmt.Printf("Int value after increment: %v\n", intValue)

	var uintValue uint16 = 65535 // Max value for uint16
	fmt.Printf("Uint value: %v\n", uintValue)
	uintValue++
	fmt.Printf("Uint value after increment: %v\n", uintValue)
}
