// Package twofer implements two for one: one for you and one for me.
package twofer

import "fmt"

// ShareWith returns the "one for you, one for me" line corresponding to the given name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
