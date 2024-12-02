package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

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
// func main() {
// 	//usign this we can append more into the array
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[0:1])
// 	prices[1] = 9.99

// 	// adding new prices using append
// 	updatedPrices := append(prices, 5.99)
// 	fmt.Println(updatedPrices, prices)

// 	//re assign new slice to prices
// 	prices = append(prices, 5.99)
// 	fmt.Println(updatedPrices, prices)
// }

func main() {
	//array with max lenght
	hobbies := [3]string{"Coding", "music", "playing"}
	fmt.Println(hobbies)

	//access first element
	fmt.Println(hobbies[0])
	//access second and third  element
	futureHobbies := hobbies[1:3]
	fmt.Println(futureHobbies)
	//access first and second  element
	fmt.Println(hobbies[0:2])

	//dynamic array with no lenght
	goal := []string{"master-go", "backend"}
	fmt.Println(goal)

	//change the last array
	goal[1] = "backend-dev"

	//add brand new third goal
	goal = append(goal, "learning all basic")
	fmt.Println(goal)

	//creating dynamic array
	products := []Product{
		{
			"first-product",
			"A first product",
			12.99,
		},
		{
			"second-product",
			"A second product",
			132.99,
		},
	}

	fmt.Println(products)

	//add brand new third product
	newProduct := Product{
		"third-product",
		"A third product",
		122.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
