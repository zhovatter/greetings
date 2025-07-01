package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
// Functions that start with a capital letter are known as exported functions
// and can be called by a function not in the same package.
func Hello(name string) (string, error) {

	//if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	//If a name was received, return a greeting that embeds the name in a message.

	// := initalizes and assigns, and is shortcut for:
	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat()) //testing if tests work when name param is not given
	return message, nil //nil represents the lack of an error
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	//Loop through the received slice of names, calling
	// the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map, associate the retrieved messages with the name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random
// randomFormat is not exported, since func name starts with lowercase letter (can only be used by functions within its own module)
func randomFormat() string {
	// A slice of messge formats.

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	//Return a randomly selected message format by specifying
	// a random index for the slice of formats

	return formats[rand.Intn(len(formats))]
}
