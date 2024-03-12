package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil avilable\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var fistName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&fistName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter your ticket: ")
		fmt.Scan(&userTickets)

		isValidName := len(fistName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, fistName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v \n",
				fistName, lastName, userTickets, email)

			fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			lastNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
				lastNames = append(lastNames, names[1])

			}
			fmt.Printf("The first name of bookings are: %v\n", firstNames)
			fmt.Printf("The last name of bookings are: %v\n", lastNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email is not contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Printf("Your input data is invalid. try again\n")
		}

	}

}
