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

	names := []string{"Xindi", "Justin", "Casper", "Arby"}

	messages, err := greetings.Hellos(names)

	// Handle error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
