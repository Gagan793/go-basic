package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex // to protect shared vars in concurrent calls

func SetupRoutes() {
	http.HandleFunc("/api/book", bookHandler)
	http.HandleFunc("/api/tickets", ticketsHandler)
}

// Struct to parse booking requests from frontend
type BookingRequest struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	NumberOfTickets uint   `json:"numberOfTickets"`
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var bookingReq BookingRequest
	err := json.NewDecoder(r.Body).Decode(&bookingReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate user data using your existing validation function
	isValidName, isValidEmail, isValidTicketCount := userValidation(
		bookingReq.FirstName, bookingReq.LastName, bookingReq.Email, bookingReq.NumberOfTickets, remainingTickets,
	)
	if !isValidName || !isValidEmail || !isValidTicketCount {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if bookingReq.NumberOfTickets > remainingTickets {
		http.Error(w, "Not enough tickets available", http.StatusBadRequest)
		return
	}

	// Use your existing function to book tickets & update globals
	bookTicket(bookingReq.NumberOfTickets, bookingReq.FirstName, bookingReq.LastName, bookingReq.Email)

	// Optionally call sendTicket (maybe async)
	go sendTicket(bookingReq.NumberOfTickets, bookingReq.FirstName, bookingReq.LastName, bookingReq.Email)

	resp := map[string]interface{}{
		"message":          fmt.Sprintf("Thanks %s %s for booking %d tickets", bookingReq.FirstName, bookingReq.LastName, bookingReq.NumberOfTickets),
		"remainingTickets": remainingTickets,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
func ticketsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	resp := map[string]interface{}{
		"conferenceName":   conferenceName,
		"remainingTickets": remainingTickets,
		"bookingsCount":    len(bookings),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
