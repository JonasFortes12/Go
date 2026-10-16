package cap9exerciseslevel04

import (
	"fmt"
)

func Exerc09() {

	var peopleMap map[string][]string = make(map[string][]string, 3)

	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			peopleMap["João_Silva"] = []string{"paintball", "fishing", "games"}
		case 1:

			peopleMap["Leticia_Soares"] = []string{"Embroidery", "Leisure", "Reading"}

		case 2:
			peopleMap["Milene_Castro"] = []string{"Trinkets", "Baking", "Origami"}

		}
	}

	/* for i, value := range peopleMap {
		fmt.Printf("Index: %v | Value: %v \n", i, value)
	} */

	peopleMap["Jonas"] = []string{"Coding", "Reading", "BJJ"}

	for i, values := range peopleMap {
		fmt.Printf("Index: %v \n Values: ", i)

		for j, value := range values {
			fmt.Printf("%v-%v ", j+1, value)
		}
		fmt.Println("")
	}

}
