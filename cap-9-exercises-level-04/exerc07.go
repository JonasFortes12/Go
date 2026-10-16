package cap9exerciseslevel04

import "fmt"

func Exerc07() {

	arr := make([][]string, 3)

	for i := 0; i < len(arr); i++ {
		switch i {
		case 0:
			arr[i] = append(arr[i], "João", "Silva", "PaintBool")
		case 1:
			arr[i] = append(arr[i], "Leticia", "Soares", "Dance")

		case 2:
			arr[i] = append(arr[i], "Milene", "Castro", "Running")

		}
	}

	for i, people := range arr {

		fmt.Printf("Index: %v | People: %v \n", i, people)
	}

}
