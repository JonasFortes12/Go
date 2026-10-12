package cap2variablesvaluestypes

import (
	"fmt"
	"os"
)

func FmtMethods() {

	//Interpreted String Literal
	s := "Hello, good morning,\n how are you ? \t I wish all Ok"

	//Raw String Literal
	s2 := `Hello, good morning,\n how are you ? \t I wish all Ok`

	fmt.Println(s)

	fmt.Println(s2)

	//-------------Kinds of Print--------------------------------

	// This method prints the string without a new line at the end
	fmt.Print("Hello, good morning,\n how are you ? \t I wish all Ok")

	// This method prints the string with a new line at the end
	fmt.Println("Hello, good morning,\n how are you ? \t I wish all Ok")

	// This method prints the string with a new line at the end and also returns the number of bytes written and an error if any
	bytes, err := fmt.Println("Hello, good morning,\n how are you ? \t I wish all Ok")
	println(bytes, err)

	// This method returns the combined string without printing it
	s3 := fmt.Sprint("Hello", " good morning", ", how are you ? ", " I wish all Ok")
	fmt.Println(s3)

	// This method returns the combined string without printing it and also returns the number of bytes written and an error if any
	s4, err := fmt.Fprint(os.Stdout, "Hello", " good morning", ", how are you ? ", " I wish all Ok")
	println(s4, err)

}
