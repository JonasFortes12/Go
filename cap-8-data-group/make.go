package cap8datagroup

func Make() {

	//Creating a slice of strings with the make function.
	planets := make([]string, 5, 10)

	// The difference between the length and the capacity of a slice is that the length is the number of elements
	// that are currently in the slice,
	// while the capacity is the number of elements that can be stored
	// in the slice before it needs to be resized.

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Mars"
	planets[4] = "Jupiter"
	planets[5] = "Saturn"
	planets[6] = "Uranus"
	planets[7] = "Neptune"

}
