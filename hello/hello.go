package main

import (
    "fmt"

    "github.com/adidassg11/greetings"
    "github.com/adidassg11/hello/morestrings"
    "github.com/google/go-cmp/cmp"

)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))

    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
