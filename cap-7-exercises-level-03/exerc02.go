package cap7exerciseslevel03

import "fmt"

func Exerc02() {

	for i := 65; i <= 90; i++ {
		fmt.Printf("Value: %v | Unicode: %v \n", string(rune(i)), i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("%#U \n", rune(i))
		}
	}

}
