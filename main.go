package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Reference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// remainingTickets := 50

	fmt.Printf("Welcome to %v Booking Application \n", conferenceName)
	fmt.Printf("We have total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket to attend")

	// var bookings []string
	bookings := []string{}

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	// ask for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets you want? ")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

	// ====== Array Informations ======
	// fmt.Printf("The whole array: %v \n", bookings)
	// fmt.Printf("The first element array: %v \n", bookings[0])
	// fmt.Printf("Array type: %T \n", bookings)
	// fmt.Printf("Array length: %v \n", len(bookings))

	fmt.Printf("These are all our bookings: %v \n", bookings)
}
