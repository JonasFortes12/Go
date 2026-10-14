package cap7exerciseslevel03

import (
	"fmt"
	"time"
)

func Exerc03() {

	currentYear := time.Now().Year()
	birthYear := 2000

	for i := birthYear; i <= currentYear; i++ {
		fmt.Printf("%v\n", i)
	}
}
