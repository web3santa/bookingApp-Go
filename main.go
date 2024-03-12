package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// userInput
	fistName, lastName, email, userTickets := getUserInput()

	// valid logic
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(fistName, lastName, userTickets, email, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		// booking
		getBookingTicket(userTickets, fistName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, fistName, lastName, email)

		fullNames := getFullNames()
		for _, fullName := range fullNames {
			fmt.Printf("The first name of bookings are: %v\n", fullName.FistNames)
			fmt.Printf("The last name of bookings are: %v\n", fullName.LastNames)
		}

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year")
			// break
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

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil avilable\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

type FullName struct {
	FistNames string
	LastNames string
}

func getFullNames() []FullName {
	var fullNames []FullName
	for _, booking := range bookings {

		fullName := FullName{
			FistNames: booking.firstName,
			LastNames: booking.lastName,
		}
		fullNames = append(fullNames, fullName)

	}
	return fullNames
}

func getUserInput() (string, string, string, uint) {
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

	return fistName, lastName, email, userTickets
}

func getBookingTicket(userTickets uint, fistName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       fistName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v \n",
		fistName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(15 * time.Second)
	var ticekt = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("sending ticket:\n %v to email address %v\n", ticekt, email)
	fmt.Println("############")
	wg.Done()
}
