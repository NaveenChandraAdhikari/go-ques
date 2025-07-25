package main
/*

Define interface Notifier with Notify(msg string) error


           Implement two types:
EmailNotifier (prints “Sending email: <msg>”)


SlackNotifier (prints “Sending slack message: <msg>”)
    Write a function SendWelcome(notifier Notifier)
 */
//anniterface defines a set of methods a type must have
// any type with a Notify(msg string) error method implements the Nptifier interface


package main

import (
"fmt"
)

// 1: Define the interface
type Notifier interface {
	Notify(msg string) error
}

// 2: EmailNotifier implementation
type EmailNotifier struct{}

func (e EmailNotifier) Notify(msg string) error {
	fmt.Printf("Sending email: %s\n", msg)
	return nil
}

// 3: SlackNotifier implementation
type SlackNotifier struct{}

func (s SlackNotifier) Notify(msg string) error {
	fmt.Printf("Sending slack message: %s\n", msg)
	return nil
}

// 4: sen d function that uses the interface
func SendWelcome(notifier Notifier) {
	err := notifier.Notify("Welcome!")
	if err != nil {
		fmt.Println("Notification error:", err)
	}
}

func main() {
	email := EmailNotifier{}
	slack := SlackNotifier{}

	SendWelcome(email)
	SendWelcome(slack)
}




/*
General Note :
// as we have also used with error
In Go, errors are represented by the error interface type. A function that can fail (like Notify) typically returns an error value.
Convention:
If a function succeeds, it returns err == nil (no error).
If a function fails, it returns err != nil (an error occurred, and err contains details about the failure).


A function like func add(a, b int) int that performs simple arithmetic might not return an error because it’s unlikely to fail.
Structs or variables don’t inherently involve error unless explicitly designed to (e.g., a struct field or custom error type).
But very common: Error handling with the error interface is a core part of Go’s design, especially in:
Functions that interact with external resources (files, networks, databases).
Interfaces defining methods that can fail.
Custom error types for specific use cases.
 */