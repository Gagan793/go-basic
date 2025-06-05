package main

import (
	"fmt"
)

//	func bookTicket(userTickets uint, firstName string, lastName string, email string) {
//		remainingTickets = remainingTickets - userTickets
//		//create a map
//		var userData = UserData{
//			firstName:       firstName,
//			lastName:        lastName,
//			email:           email,
//			numberofTickets: userTickets,
//		}
//		bookings = append(bookings, userData)
//		fmt.Printf("List of Bookings is %v\n", bookings)
//		fmt.Printf("The whole slice: %v\n", bookings)
//		fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a Confirmation email at %v \n", firstName, lastName, userTickets, email)
//		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
//	}
func bookTicket(userTickets uint, firstName, lastName, email string) error {
	if userTickets > remainingTickets {
		return fmt.Errorf("not enough tickets remaining")
	}
	remainingTickets -= userTickets
	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}
	bookings = append(bookings, userData)
	// You can log for debugging
	fmt.Printf("Booking success: %+v\n", userData)
	return nil
}
