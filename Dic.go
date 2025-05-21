package main

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := map[string]string{
		"gopher": "A person who programs in Go.",
		"goroutine": "A lightweight thread managed by the Go runtime.",
		"channel": "A way for goroutines to communicate with each other.",
	}

	var word string
	fmt.Print("Enter a word to look up: ")
	fmt.Scanln(&word)
	word = strings.ToLower(word)

	definition, ok := dictionary[word]
	if ok {
		fmt.Printf("Definition of %s: %s\n", word, definition)
	} else {
		fmt.Printf("Sorry, the word '%s' was not found in the dictionary.\n", word)
	}
}
