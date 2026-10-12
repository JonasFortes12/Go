package cap4fundamentals

import "fmt"

func Strings() {

	s := "Hello, World! ç"

	fmt.Printf("String: %v\n", s)

	sBytes := []byte(s)

	fmt.Printf("Bytes: %v\n", sBytes)

	// I will print the bytes of the string, and the string representation of the bytes.
	for _, b := range sBytes {
		fmt.Printf("%b - %T - %#U - %#x - %v\n", b, b, b, b, string(b))
	}

	fmt.Println("")

	// I will print the runes of the string, and the string representation of the runes.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x - %v\n", s[i], s[i], s[i], s[i], string(s[i]))
	}

}
