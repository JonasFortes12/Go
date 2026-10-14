package cap6ControlFlow

import "fmt"

func BreakContinue() {

	num := 0
	count := 0

	for {
		num++
		count++

		if num%2 == 0 {
			fmt.Println(num, "is Odd")
			continue
		}

		fmt.Println(num, "is Even")

		if count >= 100{
			break
		}

	}

}
