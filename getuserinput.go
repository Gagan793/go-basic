package main

import (
	"fmt"
)

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your Last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter Your Email Address")
	fmt.Scan(&email)
	fmt.Println("Enter Number of Tickets")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
