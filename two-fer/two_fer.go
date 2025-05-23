package twofer

import "fmt"

// Returns a sentence containing the name of the person
// if presents, or with "you".
func ShareWith(name string) string {
	if name == "" {
		return fmt.Sprintf("One for you, one for me.")
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
