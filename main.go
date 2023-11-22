package main

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	conferenceTickets = 50
	conferenceName    = "Go Conference"
)

var remainingTickets uint = conferenceTickets
var bookings = []string{}

func main() {

	greetUser()

	for remainingTickets > 0 && len(bookings) < conferenceTickets {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidTicketNumber, isValidEmail := validateUserInput(firstName, lastName, userTickets, email)

		if isValidTicketNumber && isValidName && isValidEmail {
			bookTicket(firstName, lastName, userTickets, email)
			if remainingTickets == 0 {
				fmt.Printf("Our %s conference is booked out. Come back next year!\n", conferenceName)
				break
			}
		} else {
			printValidationErrors(isValidEmail, isValidTicketNumber, isValidName)
			continue
		}
	}

	printBookingsByFirstNames(bookings)
}

func bookTicket(firstName, lastName string, userTickets uint, email string) {
	bookings = append(bookings, fmt.Sprintf("%s %s", firstName, lastName))
	remainingTickets -= userTickets

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
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


func validateUserInput(firstName, lastName string, userTickets uint, email string) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidEmail := isValidEmail(email)

	return isValidName, isValidTicketNumber, isValidEmail
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

func printBookingsByFirstNames(bookings []string) {
	firstNames := make([]string, len(bookings))
	for i, booking := range bookings {
		names := strings.Fields(booking)
		firstNames[i] = names[0]
	}
	fmt.Printf("These are all bookings: %v\n", firstNames)
}

func greetUser() {
	fmt.Printf("Conference tickets: %d, conference name: %s\n", conferenceTickets, conferenceName)
	fmt.Printf("Welcome to our %s booking app!\n", conferenceName)
	fmt.Printf("We have a total of %d tickets, and %d tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

