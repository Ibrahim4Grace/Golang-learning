package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)

	// numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15)

	fmt.Println(sum)
}

func factorial(number int) int {

	//recursion approach
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	//loop approach
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }
	// return result
}

// factorial = 5 * 4 * 3 * 2 * 1 => 120

//variadic function we can use ... instead of []int
func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum +val
	}
	return sum
}
