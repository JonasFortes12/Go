package cap7exerciseslevel03

import "fmt"

func Exerc06() {

	name := "Jonas"
	age := 78

	isNameCorrect := name == "Jonas"
	isOld := age > 50

	if !isNameCorrect {
		fmt.Println("Inform the correct name!")
	} else if isOld {
		fmt.Println("Pay attention on stairs")
	} else {
		fmt.Println("Welcome!")
	}
}
