package main

import (
	"fmt"
	"log"

	"0.3.1_create_module/greeting"
)

func main() {
	log.SetPrefix("greeting:")
	log.SetFlags(11) //
	message, err := greeting.Hello("")

	if err != nil {
		log.Fatal(err)
		// auto return
	}

	fmt.Println(message)
}

/* https://go.dev/doc/tutorial/call-module-code */
