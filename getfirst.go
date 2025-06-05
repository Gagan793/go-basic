package main

func getFirst() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
	//fmt.Printf("All bookings:%v\n", firstNames)
}

// func getFirstNamesHandler(w http.ResponseWriter, r *http.Request) {
// 	firstNames := getFirst()

// 	w.Header().Set("Content-Type", "application/json")
// 	err := json.NewEncoder(w).Encode(firstNames)
// 	if err != nil {
// 		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 	}
// }
// func SetupRoutes() {
// 	http.HandleFunc("/api/book", bookHandler)
// 	http.HandleFunc("/api/first-names", getFirstNamesHandler)
// 	// other routes
// }
