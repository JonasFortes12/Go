package cap8datagroup

import "fmt"

func Map() {

	contacts := make(map[string]int)

	contacts["João"] = 123456789
	contacts["Maria"] = 987654321

	//To access the elements of a map, we can use the key of the map.
	joaoContact := contacts["João"]
	mariaContact := contacts["Maria"]

	fmt.Printf("\n\nContacts: %v", contacts)
	fmt.Printf("\nJoão's contact: %v | Maria's contact: %v", joaoContact, mariaContact)

	luisContact, ok := contacts["Luis"]

	if !ok {
		fmt.Printf("\nLuis's contact not found")
	} else {
		fmt.Printf("\nLuis's contact: %v", luisContact)
	}

	//We can add new elements to a map by using the key of the map.
	contacts["Luis"] = 111222333

	//To delete an element from a map, we can use the delete function.
	delete(contacts, "Maria")

	fmt.Printf("\n\nContacts after deleting Maria: %v", contacts)

	//To iterate over a map, we can use the range keyword.
	for name, contact := range contacts {
		fmt.Printf("\nName: %v | Contact: %v", name, contact)
	}

}
