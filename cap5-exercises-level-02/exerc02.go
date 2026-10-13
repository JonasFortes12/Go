package cap5exerciseslevel02

import "fmt"

func Exerc02() {

	name := "jonas"
	age := 25

	isNameCorrect := name == "Jonas"
	isNameDefined := name != ""
	isAgeValid := !(age <= 0)
	isAdult := age >= 18
	isChild := age < 12
	isOld := age > 50

	fmt.Printf(`
		Name: %s
		Age: %d
		Is Name Correct: %t
		Is Name Defined: %t
		Is Age Valid: %t
		Is Adult: %t
		Is Child: %t
		Is Old: %t
	`, name, age, isNameCorrect, isNameDefined, isAgeValid, isAdult, isChild, isOld)

}
