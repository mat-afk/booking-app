package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Gopher Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining for the event.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n\n")

	for {
		
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

		nameIsValid := len(firstName) >= 2 && len(lastName) >= 2 
		emailIsValid := strings.Contains(email, "@")
		ticketNumberIsValid := userTickets > 0 && userTickets <= remainingTickets

		if nameIsValid && emailIsValid && ticketNumberIsValid {

			remainingTickets -= userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("\nThank you %v %v for booking %v tickets. A confirmation e-mail will be sent to %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are %v tickets remaining for the %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("\nThe first names of bookings are: %v.\n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("\nSorry, the %v is booked out. Come back next year :]\n", conferenceName)
				break
			}

		} else {
			errorMessages := []string{}

			errorMessages = append(errorMessages, "\nThere was an error with your booking:")

			if !nameIsValid {
				errorMessages = append(errorMessages, "Your first name or last name is too short.")
			}

			if !emailIsValid {
				errorMessages = append(errorMessages, "Your e-mail address doesn't contain '@' sign.")
			}

			if !ticketNumberIsValid {
				errorMessages = append(errorMessages, "Your ticket number is invalid.")
			}

			errorMessages = append(errorMessages, "Please try again.\n\n")

			fmt.Print(strings.Join(errorMessages, "\n"))
		}
	}
}
