package main

import (
    "fmt"
    "log"

    "github.com/adidassg11/greetings"
    "github.com/adidassg11/hello/morestrings"
    "github.com/google/go-cmp/cmp"

)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))

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
}
