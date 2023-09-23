package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // if no name was given, return an error with a message
    if name == "" {
        return "", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    //message := fmt.Sprintf("Hi, %v. Welcome!", name)
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// randomFormat returns one of a set of greeting messages
// returned message is selected at random

func randomFormat() string {
    // a slice of message formats
    formats := []string{
        "Hi, %v, Welcome!",
        "great to see you %v!",
        "Hail %v, well met!",
    }

    // return a randome one
    return formats[rand.Intn(len(formats))]
}

