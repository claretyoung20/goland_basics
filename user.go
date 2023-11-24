package main

// UserData represents user data with JSON tags.
type UserData struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	NumberOfTickets uint   `json:"numberOfTickets"`
}
