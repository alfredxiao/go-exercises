package main

import (
	"fmt"

	"github.com/alfredxiao/go-exercises/modules/mod_b_ref_a/greetings/greetings"
	"github.com/alfredxiao/go-exercises/modules/mod_b_ref_a/greetings/utils"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

	fmt.Println(utils.Utils())
}
