package main

import (
	"fmt"
	"strings"
	"time"
	"sync"
)

const conferenceTickets uint = 50
const conferenceName string = "Gopher Conference"
var remainingTickets uint = conferenceTickets
var bookings = []User{}

type User struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()


	firstName, lastName, email, userTickets := getUserData()

	nameIsValid, emailIsValid, ticketNumberIsValid := validateUserData(firstName, lastName, email, userTickets, remainingTickets)

	if nameIsValid && emailIsValid && ticketNumberIsValid {

		bookTickets(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		fmt.Printf("\nThe first names of bookings are: %v.\n\n", getFirstNames())

		if remainingTickets == 0 {
			fmt.Printf("\nSorry, the %v is booked out. Come back next year :]\n", conferenceName)
			// break
		}

	} else {
		errorMessages := []string{}

		errorMessages = append(errorMessages, "\nThere was an error with your booking:")

		if !nameIsValid {
			errorMessages = append(errorMessages, "Your first name or last name is too short.")
		}

		if !emailIsValid {
			errorMessages = append(errorMessages, "Your e-mail address doesn't contain the '@' sign.")
		}

		if !ticketNumberIsValid {
			errorMessages = append(errorMessages, "Your ticket number is invalid.")
		}

		errorMessages = append(errorMessages, "Please try again.\n\n")

		fmt.Print(strings.Join(errorMessages, "\n"))
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still remaining for the event.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserData() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	var userData = User {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	bookings = append(bookings, userData)

	fmt.Printf("\nThank you %v %v for booking %v tickets. A confirmation e-mail will be sent to %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for the %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n###############")
	fmt.Println("Sending ticket...")
	fmt.Printf("\n%v \nto email address %v.\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
