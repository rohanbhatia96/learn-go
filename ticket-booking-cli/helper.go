package main

import (
	"fmt"
	"time"
)

func greetUsers(confName string) {
	fmt.Printf("Welcome to %v booking application\n", confName)
}

func printBookingDetails(bookings []UserData) {
	var displayStrings []string
	for _, booking := range bookings {
		displayStrings = append(displayStrings, "Buyer: "+booking.name)
	}
	fmt.Printf("All Bookings: %v\n", displayStrings)
}

func sendTickets(user UserData) {
	var ticket = fmt.Sprintf("%v ticket for %v sent to %v", user.tickets, user.name, user.email)
	time.Sleep(10 * time.Second)
	fmt.Println("############")
	fmt.Println("############")
	fmt.Println(ticket)
	wg.Done()
}
