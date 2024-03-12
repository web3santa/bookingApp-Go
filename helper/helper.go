package helper

import "strings"

func ValidateUserInput(fistName string, lastName string, userTickets uint, email string, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(fistName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
