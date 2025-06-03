package main

func getFirst() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
	//fmt.Printf("All bookings:%v\n", firstNames)
}
