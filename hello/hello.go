package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message - NO ERROR
	message, err := greetings.Hello("Wagner")
	fmt.Println(message)

	// A slice of names.
	names := []string{"Alana", "Samantha", "Bruna"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

	// Request a greeting message.
	messageerror, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messageerror)
}
