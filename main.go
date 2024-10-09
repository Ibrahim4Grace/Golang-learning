package main

import "fmt"

func main() {

    var conferenceName = "Conference"
    const ConferenceTickets = 50
    var remainingTickets = 30

    fmt.Printf("Welcom to our %v booking application\n", conferenceName)
    fmt.Printf("we've total of %v tickets and %v are still availabe...\n", ConferenceTickets, remainingTickets)
    fmt.Println("Get your ticket here to attend...")

    // var conferenceName = "Go Conference"
    // fmt.Println (conferenceName)


    var firstName string 
    var lastName string 
    var email string 
    var userTickets uint
    // ask user for their name
    //& stands for pointer to user
    // fmt.Scan(&userName)

    fmt.Println("Enter your first name...")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name...")
    fmt.Scan(&lastName)

    fmt.Println("Enter your email...")
    fmt.Scan(&email)

    fmt.Println("Enter how mnay ticket you want...")
    fmt.Scan(&userTickets)

    remainingTickets = remainingTickets - userTickets

    fmt.Printf("Thnak you %v %v for booking %v tickets. You will receive email confirmation at %v\n", firstName,lastName,email, userTickets)

    fmt.Printf(" %v ticket remaing for  %v \n", remainingTickets, userTickets)

}

