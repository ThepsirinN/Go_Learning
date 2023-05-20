package main

import (
	"fmt"
	"log"

	"0.3.1_create_module/greeting"
)

func main() {
	log.SetPrefix("greeting:")
	log.SetFlags(11) //
	// message, err := greeting.Hello("Barko")
	names := []string{"barko", "bank", "boom", "bell", "eiei"}
	message, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err)
		// auto return
	}

	// fmt.Println(message)
	fmt.Println(message)
}
