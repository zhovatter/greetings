package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// Functions that start with a capital letter are known as exported functions
// and can be called by a function not in the same package.
func Hello(name string) (string, error) {

	//if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	//If a name was receives, return a greeting that embeds the name in a message.

	// := initalizes and assigns, and is shortcut for:
	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil //nil represents the lack of an error
}
