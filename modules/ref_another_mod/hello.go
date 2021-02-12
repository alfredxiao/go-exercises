package main

import (
	"fmt"

	"github.com/alfredxiao/go-demo-module/greetings"
	v2 "github.com/alfredxiao/go-demo-module/v2/greetings"
)

func main() {
	// Get a greeting message and print it.

	messageV1 := greetings.Hello("Gladys")
	fmt.Println(messageV1)

	messageV2, err := v2.Hello("Oldman")
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Printf(messageV2)
}
