package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil avilable\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets
	bookings[0] = fistName + " " + lastName

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value : %v\n", bookings[0])
	fmt.Printf("The Type : %T\n", bookings)
	fmt.Printf("The array elngth : %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v \n",
		fistName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

}
