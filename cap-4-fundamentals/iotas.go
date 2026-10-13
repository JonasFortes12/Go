package cap4fundamentals

import "fmt"

func IotasAndBitDislocations() {

	// Iota is a predeclared identifier representing the untyped integer ordinal number
	// of the constant in a const declaration.
	// It is reset to 0 whenever the reserved word const appears
	// in the source and increments after each ConstSpec.
	// It can be used to construct a set of related constants that are easy to maintain.
	const (
		a = iota // 0
		b        // 1
		c        // 2
	)

	// Bits dislocations are used to shift the bits of a number to the left or right.

	x := 20

	y := x << 1 // Shift the bits of x to the left by 1 position, equivalent to multiplying by 2

	z := x >> 1 // Shift the bits of x to the right by 1 position, equivalent to dividing by 2

	fmt.Printf("Bits of x: %b | Value %v\n", x, x) // 1 1

	fmt.Printf("Bits of y: %b | Value %v\n", y, y) // 10 2

	fmt.Printf("Bits of z: %b | Value %v\n", z, z) // 0 0
}
