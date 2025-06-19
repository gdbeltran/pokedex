package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleaned[0])
	}
}

func cleanInput(text string) []string {
	var words []string
	lowerWords := strings.ToLower(text)
	splitWords := strings.Split(lowerWords, " ")
	for _, word := range splitWords {
		cleanedWord := strings.TrimSpace(word)
		if cleanedWord != "" {
			words = append(words, cleanedWord)
		}
	}
	return words
}
