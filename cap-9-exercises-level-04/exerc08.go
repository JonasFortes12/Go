package cap9exerciseslevel04

import "fmt"

func Exerc08() {

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

	for i, value := range peopleMap {
		fmt.Printf("Index: %v | Value: %v \n", i, value)
	}

}
