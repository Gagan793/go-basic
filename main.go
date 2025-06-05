package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Initialize routes
	SetupRoutes()

	// Enable CORS
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start server with CORS
	http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods)(http.DefaultServeMux))
}

// package main

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// this is empty list of maps
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

// // func main() {

// // 	//fmt.Println("Welcome to", conferenceName, "application")
// // 	//fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
// // 	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

// // 	greetUsers()
// // 	for {
// // 		firstName, lastName, email, userTickets := getUserInput()

// // 		isValidName, isValidEmail, isValidTicketCount := userValidation(firstName, lastName, email, userTickets, remainingTickets)
// // 		if isValidEmail && isValidName && isValidTicketCount {
// // 			//fmt.Printf("We onle have %v tickets\n", remainingTickets)
// // 			bookTicket(userTickets, firstName, lastName, email)
// // 			go sendTicket(userTickets, firstName, lastName, email)
// // 			firstNames := getFirst()
// // 			fmt.Printf("the first names of bookings are: %v\n", firstNames)
// // 			if remainingTickets == 0 {
// // 				fmt.Println("Conference is Sold out")
// // 				break
// // 			}

// // 		} else {
// // 			fmt.Println("Invalid Credentials")
// // 			break
// // 		}

// // 	}

// // }

// var conferenceName = "Go Conference"
// var remainingTickets uint = 50

// var bookings = make([]UserData, 0)

// type UserData struct {
// 	firstName       string
// 	lastName        string
// 	email           string
// 	numberofTickets uint
// }
