package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting messaage
	message, err := greetings.Hello("Andy")
	// If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}
	
	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	// If no error was returned, print the returned message to the console
	fmt.Println(message)
}
