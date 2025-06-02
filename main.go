package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Println("Welcome to", conferenceName, "application")
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Book your tickets here")
	var bookings []string

	for {
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
		if userTickets > remainingTickets {
			fmt.Printf("We onle have %v tickets\n", remainingTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a Confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("All bookings:%v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Println("Conference is Sold out")
			break
		}
	}
}

// making a booking application
// step 1 - welcome the users
