package dict

import (
	"errors"
	"fmt"
)

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Word Not found. Can't Update the word")
	errNotDelete  = errors.New("Not found the word to delete")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	fmt.Println("\nSearch [", word, "] ...")
	value, exists := d[word]
	if exists {
		fmt.Println("found  :", word)
		return value, nil
	}
	return "", errNotFound
}

// Add
func (d Dictionary) Add(word, def string) error {
	fmt.Println("\n  Add  : ", word)
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, newWord string) error {
	fmt.Println("\n  Update definiton :", word, " to ", newWord)
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newWord
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	fmt.Println("\n  Delete a word", word)
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errNotDelete
	}
	return nil
}
