package main

import "fmt"

func main() {

	conferenceName := "Gopher Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	// slice: dynamic array
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining for the event.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n\n")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your e-mail address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to book? ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("\nThank you %v for booking %v tickets. A confirmation e-mail will be sent to %v.\n", bookings[0], userTickets, email)
	fmt.Printf("There are %v tickets remaining for the %v.\n", remainingTickets, conferenceName)

	fmt.Printf("\nThese are all of our bookings %v.\n", bookings)
}
