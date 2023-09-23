package main

import (
	"fmt"
	"log"

	"github.com/adidassg11/greetings"
	"github.com/adidassg11/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/example/hello/reverse"
)

func split(sum int) (x, y int) {
	// Grabbed this from https://go.dev/tour/basics/7
	x = sum * 4 / 9
	y = sum - x
	// Only use naked returns in short functions
	return // "naked" return because we've already named the retuns in the func definition
}

// Used `go work use ./example/hello` from https://go.dev/doc/tutorial/workspaces

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(split(200))

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Gladys") // this should work fine
	//message, err := greetings.Hello("") // this should error
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// if no error, print message to console
	fmt.Println(message)
	fmt.Println(reverse.String("Hello"), reverse.Int(24601)) // From https://go.dev/doc/tutorial/workspaces
}
