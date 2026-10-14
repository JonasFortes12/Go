package cap7exerciseslevel03

import "fmt"

func Exerc01() {

	for i := 1; i <= 10000; i++ {
		fmt.Printf("%v ", i)
		if i%10 == 0 {
			fmt.Print("\n")
		}
	}

}
		