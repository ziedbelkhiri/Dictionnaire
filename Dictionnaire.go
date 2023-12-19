package main

import (
	"fmt"
	"sort"
)

// Dictionary represents a mapping of words to their definitions
type Dictionary map[string]string

// Add inserts a new word and its definition into the dictionary
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Get retrieves the definition of a specific word from the dictionary
func (d Dictionary) Get(word string) string {
	return d[word]
}

// Remove deletes a word and its definition from the dictionary
func (d Dictionary) Remove(word string) {
	delete(d, word)
}

// List prints out the sorted list of words and their definitions
func (d Dictionary) List() {
	var sortedWords []string
	for word := range d {
		sortedWords = append(sortedWords, word)
	}
	sort.Strings(sortedWords)

	fmt.Println("Dictionary:")
	for _, word := range sortedWords {
		fmt.Printf("%s: %s\n", word, d[word])
	}
}

func main() {
	// Initialize an empty dictionary
	myDictionary := make(Dictionary)

	// Add words and their definitions
	myDictionary.Add("apple", "a fruit")
	myDictionary.Add("go", "a programming language")
	myDictionary.Add("house", "a place to live")

	// Get and display the definition of a specific word
	word := "go"
	definition := myDictionary.Get(word)
	fmt.Printf("Definition of '%s': %s\n", word, definition)

	// Remove a word from the dictionary
	wordToRemove := "house"
	myDictionary.Remove(wordToRemove)
	fmt.Printf("'%s' has been removed from the dictionary.\n", wordToRemove)

	// List all words and their definitions in the dictionary
	myDictionary.List()
}