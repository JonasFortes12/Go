package cap8datagroup

import "fmt"

func MultiDimensional() {

	//Creating a multi-dimensional slice of strings.
	planets := [][]string{
		{"Mercury", "Venus", "Earth", "Mars"},
		{"Jupiter", "Saturn", "Uranus", "Neptune"},
	}
	fmt.Printf("\n\nPlanets: %v", planets)

	//Creating a multi-dimensional slice of integers.

	numbers := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	fmt.Printf("\n\nNumbers: %v", numbers)

	//If i want a multi-dimensional of different types, i can use the empty interface type.
	mixed := [][]interface{}{
		{"Mercury", 1, true},
		{"Venus", 2, false},
	}

	fmt.Printf("\n\nMixed: %v", mixed)

	//To access the elements of a multi-dimensional slice,
	// we can use the index of the outer slice and the index of the inner slice.
	fmt.Printf("\n\nFirst planet: %v", planets[0][0])
	fmt.Printf("\n\nEighth planet: %v", planets[1][3])
}
