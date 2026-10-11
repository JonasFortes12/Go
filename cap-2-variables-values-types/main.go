package main

import "fmt"

func main() {

	_, err := fmt.Println("Hello World!")

	//var name string = "jonas"

	println(err)

	hasError := err != nil

	print("Has error?", hasError)
}
