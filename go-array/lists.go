package main

import "fmt"

// func main() {
// 	productName := [4]string{"ade", "john", "kay", "kim"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	// to incude item in array
// 	productName[2] = "Carpet"
// 	fmt.Println(prices)
// 	fmt.Println(productName)
// 	//to get intem in array
// 	fmt.Println(prices[2])

// 	//how to use slice-pick array
// 	featuredPrices := prices[1:3]

// 	featuredPrices[0] = 199.99
// 	fmt.Println(featuredPrices)

// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }

//dynamic lists with slices
func main() {
	//usign this we can append more into the array
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	// adding new prices using append
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)

	//re assign new slice to prices
	prices = append(prices, 5.99)
	fmt.Println(updatedPrices, prices)
}
