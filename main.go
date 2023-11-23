package main

import (
	"booking_app/helper"
	"fmt"
	"sync"
	"time"
)

const (
	conferenceTickets = 50
	conferenceName    = "Go Conference"
)

var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

var waitGroup = sync.WaitGroup{}

func main() {
	fmt.Println(helper.MyVariable)

	greetUser()

	for remainingTickets > 0 && len(bookings) < conferenceTickets {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidTicketNumber, isValidEmail := helper.ValidateUserInput(
			firstName, lastName, userTickets, email, remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail {

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}

			bookTicket(userData)

			waitGroup.Add(1)
			go sendTicket(userData)

			if remainingTickets == 0 {
				fmt.Printf("Our %s conference is booked out. Come back next year!\n", conferenceName)
				break
			}
		} else {
			printValidationErrors(isValidEmail, isValidTicketNumber, isValidName)
			continue
		}
	}

	printBookings(bookings)

	waitGroup.Wait()
}

func bookTicket(userData UserData) {

	bookings = append(bookings, userData)
	remainingTickets -= userData.numberOfTickets

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", 
	userData.firstName, userData.lastName, userData.numberOfTickets, userData.email)
	fmt.Printf("%d tickets are left for %s\n", remainingTickets, conferenceName)
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string

	userTickets := getTicketNumber()

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	return firstName, lastName, email, userTickets
}
func getTicketNumber() uint {
	var userTickets uint
	for {
		fmt.Print("Enter the number of tickets: ")
		if _, err := fmt.Scan(&userTickets); err == nil && userTickets > 0 {
			break
		}
		fmt.Println("Invalid input. Please enter a valid positive integer.")
		fmt.Scanln()
	}
	return userTickets
}

func printValidationErrors(isValidEmail, isValidTicketNumber, isValidName bool) {
	if !isValidEmail {
		fmt.Println("Please provide a valid email.")
	}
	if !isValidTicketNumber {
		fmt.Printf("We have only %d tickets left. Please provide a valid ticket number.\n", remainingTickets)
	}
	if !isValidName {
		fmt.Println("Please provide a valid first name and last name; valid names should be at least 2 characters long.")
	}
}

func printBookings(bookings []UserData) {
	if len(bookings) == 0 {
		fmt.Println("There are no bookings.")
		return
	}
	fmt.Printf("These are all bookings: %v \n\n", bookings)
}

func greetUser() {
	fmt.Printf("Conference tickets: %d, conference name: %s\n", conferenceTickets, conferenceName)
	fmt.Printf("Welcome to our %s booking app!\n", conferenceName)
	fmt.Printf("We have a total of %d tickets, and %d tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")
}

func sendTicket(userData UserData) {
	time.Sleep(10 * time.Second) // simulate delay
	var userTicket = fmt.Sprintf("%v ticket for user: %v %v",
	userData.numberOfTickets, userData.firstName, userData.lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n to %v \n to email address: %v\n", userTicket, userData.email)
	fmt.Println("################################")
	waitGroup.Done()
}
