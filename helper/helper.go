package helper

import "regexp"

var MyVariable = "my variable"
func ValidateUserInput(firstName, lastName string, userTickets uint, email string, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidEmail := isValidEmail(email)

	return isValidName, isValidTicketNumber, isValidEmail
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}