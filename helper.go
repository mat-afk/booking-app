package main

import "strings"

func validateUserData(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	nameIsValid := len(firstName) >= 2 && len(lastName) >= 2 
	emailIsValid := strings.Contains(email, "@")
	ticketNumberIsValid := userTickets > 0 && userTickets <= remainingTickets

	return nameIsValid, emailIsValid, ticketNumberIsValid
}