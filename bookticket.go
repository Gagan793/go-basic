package main

import (
	"fmt"
	"strconv"
)

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//create a map
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	//convert userticket to string
	userData["number"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a Confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
