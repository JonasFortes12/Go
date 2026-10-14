package cap6ControlFlow

import (
	"fmt"
	"time"
)

func NestedLoop() {

	//Clock

	for hour := 0; hour < 24; hour++ {

		for min := 0; min < 60; min++ {

			for sec := 0; sec < 60; sec++ {

				time.Sleep(time.Second)
				fmt.Printf("%v:%v:%v \n", hour, min, sec)

			}

		}

	}

}
