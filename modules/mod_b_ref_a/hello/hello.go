package main

import (
	"fmt"

	g "github.com/alfredxiao/go-exercises/modules/mod_b_ref_a/greetings/greetings"
	u "github.com/alfredxiao/go-exercises/modules/mod_b_ref_a/greetings/utils"
)

func main() {
	// Get a greeting message and print it.
	message := g.Hello("Gladys")
	fmt.Println(message)

	fmt.Println(u.Utils())
}
