package main

import (
	storage "bank/storages"
	"fmt"

	"github.com/Pallinder/go-randomdata" // to use external pkg
)

//global veriable
const accountBalanceFile = "balance.txt"


func main(){

	var accountBalance, err  = storage.GetFloatFromfile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")// to set apart
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())


	//infitnite loop
	for  {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	
		if choice == 1{
			fmt.Println("Your balance is ", accountBalance)
		} else if choice == 2{
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // allow user to retry
			}
			accountBalance += depositAmount 
			storage.WriteFloatToFile(accountBalance, accountBalanceFile)
	
			fmt.Println("Balance updated! New balance:", accountBalance)
		}  else if choice == 3{
			fmt.Print("Your withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
	
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue // to retry
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You cant withdraw more than you have.")
				continue
			}
	
			accountBalance -= withdrawAmount
			storage.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New balance:", accountBalance )
		}  else  {
			fmt.Println("Goodbye" )
			break
		} 
	}

 fmt.Println("Thanks for choosing our bank")
}


