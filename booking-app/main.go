package main

import (
	"fmt"
	"sync"
	"time"
)

// Package Variable
const conferenceTickets = 50

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
	bookings              = make([]userData, 0)
)

type userData struct {
	firstName       string
	lastName        string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	// Function Call for Greeting
	greetUsers()

	for {
		// Function Call for User Input
		firstName, lastName, userTickets := returnUserInput()

		// Function Call for Validate User Input - helper.go File
		isValidName, isValidTicketNumber := validateUserInput(
			firstName,
			lastName,
			userTickets,
		)

		if isValidName && isValidTicketNumber {
			// Function Call for Booking Tickets
			bookTicket(userTickets, firstName, lastName)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName)

			// Function Call for Fetch and Print First Names
			firstNames := returnFirstNames()
			fmt.Printf("The First Names of Bookings = %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf(
					"Sorry, Our %v is Booked Out. Please Come Back Next Year\n",
					conferenceName,
				)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your Input Name is Invalid - First Name or Last Name You Entered is Short, Please Try Again")
			}
			if !isValidTicketNumber {
				fmt.Println("Your Input Number of Tickets is Invalid, Please Try Again")
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking-App\n", conferenceName)
	fmt.Printf(
		"conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n",
		conferenceName,
		conferenceTickets,
		remainingTickets,
	)
	fmt.Printf(
		"We Have Total of %v Tickets and Still We Have %v Tickets Available\n",
		conferenceTickets,
		remainingTickets,
	)
	fmt.Println("Get Your Tickets Here to Attend")
}

func returnFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func returnUserInput() (string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint

	// ask for first name and last name
	fmt.Print("Please Enter Your First Name : ")
	fmt.Scan(&firstName)
	fmt.Print("Please Enter Your Last Name : ")
	fmt.Scan(&lastName)
	fmt.Print("Please Enter Number of Tickets : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}

func bookTicket(
	userTickets uint,
	firstName string,
	lastName string,
) {
	remainingTickets = remainingTickets - userTickets

	// Create a Structure for a User
	// var userData = map[string]string{}
	// var userData map[string]string
	userData := userData{
		firstName:       firstName,
		lastName:        lastName,
		numberOfTickets: userTickets,
	}

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)

	fmt.Printf("%v %v Booked Ticket Number %v\n", firstName, lastName, userTickets)
	fmt.Printf("%v Tickets Remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("All Our Bookings = %v\n", bookings)
}

func sendTicket(userTickets uint, firstName string, lastName string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending %v for %v %v's Email\n", ticket, firstName, lastName)
	fmt.Println("########")
	wg.Done()
}
