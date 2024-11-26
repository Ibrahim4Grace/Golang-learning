package main

import (
	"errors"
	"fmt"
	"os"
)


func main(){

	revenue := getInputWithRetry("Revenue: ")
	expenses := getInputWithRetry("expenses: ")
	taxRate:= getInputWithRetry("taxRate: ")

		// we call our  function

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n",  profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	storeResults(ebt, profit,ratio)
}


func getUserInput(infoText string)(float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
    // including validation 
	if userInput <= 0 {
		return 0, errors.New("value must be greater than 0")
	}

	return userInput, nil
}

// Function to repeatedly get valid input from the user
func getInputWithRetry(prompt string) float64 {
	for {
		value, err := getUserInput(prompt)
		if err == nil {
			return value 
		}
		fmt.Println("Error:", err) 
	}
}


//how to use function in Golang 
func calculateFinancials(revenue, expenses, taxRate float64)( float64, float64, float64){
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / (1 + profit)
	return ebt, profit, ratio
}


//To store data in a file
func storeResults(ebt, profit, ratio float64) error {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %2.f\n", ebt,profit,ratio) // Formatting to 2 decimal places
	err := os.WriteFile("results.txt", []byte(results), 0644)
	if err != nil {
		return errors.New("failed to write balance to file")
	}
	return nil
}