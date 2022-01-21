package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Xindi")

	// Handle error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
