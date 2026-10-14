package cap6ControlFlow

import "fmt"

func LoopPrinting() {

	for i := 33; i <= 122; i++ {
		unicode := string(rune(i))
		fmt.Printf("Value: %v | Unicode: %v\n", i, unicode)
	}

}
