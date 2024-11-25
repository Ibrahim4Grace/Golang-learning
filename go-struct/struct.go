package main

import (
	"fmt"
	"struct/user"
)


func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//To create new User
	appUser := user.New(firstName,lastName,birthdate,)

		
    user.OutPutUserDetails( appUser)


}


func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

