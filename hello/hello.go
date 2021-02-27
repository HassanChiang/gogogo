package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message, err := greetings.Hello("Hassan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"A", "B", "C"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
