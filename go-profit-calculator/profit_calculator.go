package main

import (
	"fmt"
)
func main(){

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("expenses: ")
	taxRate := getUserInput("taxRate: ")

    // ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	// fmt.Println("Your earning before tax",  ebt)
	// fmt.Println("Your profit", profit)
	// fmt.Println("Your ratio", ratio)

		// or function

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n",  profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
	git 
}

func getUserInput(infoText string)float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

//how to use function in Golang 
func calculateFinancials(revenue, expenses, taxRate float64)( float64, float64, float64){
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / (1 + profit)
	return ebt, profit, ratio
}