package cap5exerciseslevel02

import "fmt"

func Exerc06() {

	const (
		year  = 2026
		year1 = iota + year
		year2
		year3
		year4
	)

	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

}
