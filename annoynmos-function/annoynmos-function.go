package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	double := createTransformer(2)
	tripple := createTransformer(3)

	//in here we using annoynmos function and not reuseable
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	trippled := transformNumbers(&numbers, tripple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(trippled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

//factory function  the factor will times with number
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
