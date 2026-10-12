package main

import "fmt"

type hotdog int

var h hotdog = 101

func CreatingTypes() {

	x := 10

	fmt.Printf("%v \n", x)
	fmt.Printf("%v\n", h)

	h = hotdog(x)

	fmt.Printf("%v \n", h)

}
