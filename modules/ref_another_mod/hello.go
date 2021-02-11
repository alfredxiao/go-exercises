package main

import (
	"fmt"

	"github.com/alfredxiao/go-demo-module/greetings"
	v2 "github.com/alfredxiao/go-demo-module/v2/greetings"
)

func main() {
	// Get a greeting message and print it.
	fmt.Println(greetings.Hello("Oldman"))
	message := v2.Hello("Gladys")
	fmt.Println(message)
}
