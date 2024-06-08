package main

import (
	"fmt"
	"sync"
)

type UserData struct {
	name    string
	email   string
	tickets uint
}

var wg = sync.WaitGroup{}

func main() {
	conferenceName := "Go Conference"
	greetUsers(conferenceName)

	const totalTickets uint = 50
	var remainingTickets uint = 50
	var bookings = make([]UserData, 0)

	fmt.Printf("We have a total of %v tickets and %v are still available\n", totalTickets, remainingTickets)

	for {
		var userName string
		var userEmail string
		var userTickets uint

		fmt.Print("Enter your name: ")
		n, err := fmt.Scan(&userName)
		if n != 1 || err != nil {
			fmt.Printf("There is an error in your input: %v\n", err)
			continue
		}
		fmt.Print("Enter number of tickets you want to buy: ")
		n, err = fmt.Scan(&userTickets)
		if n != 1 || err != nil {
			fmt.Printf("There is an error in your input: %v\n", err)
			continue
		}
		fmt.Print("Enter your email: ")
		n, err = fmt.Scan(&userEmail)
		if n != 1 || err != nil {
			fmt.Printf("There is an error in your input: %v\n", err)
			continue
		}

		if remainingTickets < userTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v.\n", remainingTickets, userTickets)
			continue
		}

		var userData = UserData{
			name:    userName,
			email:   userEmail,
			tickets: userTickets,
		}

		bookings = append(bookings, userData)

		wg.Add(1)
		go sendTickets(userData)

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("User %v booked %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
		printBookingDetails(bookings)
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		} else {
			fmt.Printf("Our conference is still available. %v tickets still available.\n", remainingTickets)
		}

	}
	wg.Wait()
}
