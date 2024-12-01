package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
func main() {
	userName := make([]string, 2, 5)
	userName[0] = "julie"

	userName = append(userName, "max")
	userName = append(userName, "korex")

	fmt.Println(userName)
	//using mak function in map
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.3
	courseRatings["angular"] = 4.9

	courseRatings.output()

	//using 4loops with array, slive and maps

	//to go after username with use range key word
	for index, value := range userName {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}

	//for map
	for key, value := range courseRatings {
		fmt.Println("key:", key)
		fmt.Println("value: ", value)
	}
}
