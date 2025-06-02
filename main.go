package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

// const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	//fmt.Println("Welcome to", conferenceName, "application")
	//fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketCount := userValidation(firstName, lastName, email, userTickets, remainingTickets)
		if isValidEmail && isValidName && isValidTicketCount {
			//fmt.Printf("We onle have %v tickets\n", remainingTickets)
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirst()
			fmt.Printf("the first names of bookings are: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Conference is Sold out")
				break
			}

		} else {
			fmt.Println("Invalid Credentials")
			break
		}

	}

}
func greetUsers() {
	fmt.Println("Welcome to Our conference")
	fmt.Println("Book your tickets here")
}
func getFirst() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
	//fmt.Printf("All bookings:%v\n", firstNames)
}
func userValidation(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketCount
}
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
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a Confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
