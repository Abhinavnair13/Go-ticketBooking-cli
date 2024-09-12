// 1). go mod init <project_name>
// 2). package main
// 3). import "fmt"

package main

import (
	"abhinav/booking-app-go/booking_app/helper"
	"fmt"
	"sync"
	"time"
)

// intialising a variable
var conferencesName = "Go Conferences"
var conferencesTickets int = 50
var remainingTickets uint = 50

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// waitgroup
var wg = sync.WaitGroup{}

func main() {

	greetUser()
	firstName, lastName, email, userTickets := getUserInput()

	// validateUserInput function called
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		// calling bookTicket Function
		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		// call function printFirstName
		firstNames := getFirstName()
		fmt.Printf("The first names of booking are: %v\n", firstNames)

		// if statement
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conferences is booked out, Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First Name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Printf("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Printf("Number of Tickets you entered is inValid")
		}
	}
	// // Wait: Blocks until the WaitGroup counter is 0.
	// //}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferencesName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferencesTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// parameter that use in funtion[i/p parameters   o/p parameters]
func getFirstName() []string {
	firstNames := []string{}
	// for each loop
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email id")
	fmt.Scan(&email)

	fmt.Println("Enter no of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	// as tickets are booking we need to decrease remainingTickets count
	remainingTickets = remainingTickets - uint(userTickets)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferencesName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
