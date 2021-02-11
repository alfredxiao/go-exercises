package main

import (
	"fmt"

	g "example.com/greetings"
	u "example.com/utils"
)

func main() {
	// Get a greeting message and print it.
	message := g.Hello("Gladys")
	fmt.Println(message)

	fmt.Println(u.Utils())
}
