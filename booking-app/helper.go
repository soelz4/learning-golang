package main

func validateUserInput(
	firstName string,
	lastName string,
	userTickets uint,
) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTicketNumber
}
