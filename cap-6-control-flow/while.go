package cap6ControlFlow

import (
	"fmt"
	"time"
)

func While() {

	count := 1
	for {

		if count > 10 {
			break
		}
		time.Sleep(100 * time.Millisecond)

		fmt.Println("Go Down!")
		count++
	}

}
