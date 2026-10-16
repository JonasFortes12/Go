package cap8datagroup

import "fmt"

func Slices() {

	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	/* for index, planer := range planets {
		fmt.Printf("\nIndex: %v | Planet: %v", index, planer)
	} */

	for i := 0; i < len(planets); i++ {
		fmt.Printf("\nIndex: %v | Planet: %v", i, planets[i])
	}

	fmt.Printf("\n\nAfter append Pluto: %v", planets)

	threeMoreClosePlanetsToSun := planets[0:3]
	fmt.Printf("\n\nThree more close planets to the Sun: %v", threeMoreClosePlanetsToSun)

	/* for i := 0; i < len(threeMoreClosePlanetsToSun); i++ {
		fmt.Printf("\nIndex: %v | Planet: %v", i, threeMoreClosePlanetsToSun[i])
	} */

	threeMoreFarPlanetsToSun := planets[len(planets)-3:]
	fmt.Printf("\n\nThree more far planets to the Sun: %v", threeMoreFarPlanetsToSun)

	celestialBodies := []string{"Sun", "Moon", "Pluto", "Earth"}

	//We have to spread the planets slice to append it to the celestialBodies slice.
	celestialBodies = append(celestialBodies, planets...)
	fmt.Printf("\n\nCelestial bodies: %v\n", celestialBodies)

	//Deleting the element in index 2 (Pluto) from the celestialBodies slice.

	celestialBodies = append(celestialBodies[:2], celestialBodies[3:]...)
	fmt.Printf("\n\nCelestial bodies 2: %v\n", celestialBodies)

	// There is a trick in slices:

	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Printf("\n\nSlice 1: %v", slice1)

	slice2 := append(slice1[:2], slice1[4:]...)

	fmt.Printf("\n\nSlice 1: %v | Slice 2: %v", slice1, slice2)

	// Note that the slice1 has been modified because the underlying array has been modified.
	//This is because slices are references to arrays, and when we modify a slice,
	// we are modifying the underlying array.

}
