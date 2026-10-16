package cap9exerciseslevel04

import "fmt"

func Exerc03() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	one_until_three := slice[0:3]

	five_until_last := slice[4:]

	second_until_seven := slice[1:7]

	three_until_penult := slice[2 : len(slice)-1]

	fmt.Println(one_until_three)
	fmt.Println(five_until_last)
	fmt.Println(second_until_seven)
	fmt.Println(three_until_penult)
}
 