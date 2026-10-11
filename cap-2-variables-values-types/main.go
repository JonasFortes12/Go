package main

import "fmt"

func main() {

	_, err := fmt.Println("Hello World!")

	//var name string = "jonas"

	fmt.Println(err)

	hasError := err != nil

	fmt.Println("Has error?", hasError)

	x := 10
	y := "This is  a text"

	y, z := "This is other text", 0.12

	fmt.Printf("value: %v, Type: %T \n Value: %v, Type: %T \n", x, x, y, y)

	fmt.Println(z)
	
}
