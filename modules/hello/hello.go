package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate | log.Lshortfile)
	//log.SetFlags(0)

	message, err := greetings.Hello("Eddy")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	//Call the multi name function
	names := []string{"Rocio", "Leo", "Eddy"}
	messages, error := greetings.Hellos(names)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
