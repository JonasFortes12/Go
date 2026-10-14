package cap7exerciseslevel03

import "fmt"

func Exerc08() {

	name := "Jonas"
	age := 26

	isNameCorrect := name == "Jonas"
	isOld := age > 50

	switch {
	case !isNameCorrect:
		fmt.Println("Inform the correct name!")
	case isOld:
		fmt.Println("Pay attention on stairs")
	default:
		fmt.Println("Welcome!")
	}

}
