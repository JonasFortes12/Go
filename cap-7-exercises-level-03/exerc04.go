package cap7exerciseslevel03

import (
	"fmt"
	"time"
)

func Exerc04() {

	currentYear := time.Now().Year()
	birthYear := 2000
	count := 0

	for {
		fmt.Printf("%v\n", count+birthYear)

		if count+birthYear == currentYear {
			break
		}

		count++
	}
}
