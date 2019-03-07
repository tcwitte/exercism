// Package house implements the "This is the house that Jack built" verse.
package house

import (
	"bytes"
)

const testVersion = 1

var subjects = [...]string{"house that Jack built", "malt", "rat", "cat", "dog", "cow with the crumpled horn", "maiden all forlorn", "man all tattered and torn", "priest all shaven and shorn", "rooster that crowed in the morn", "farmer sowing his corn", "horse and the hound and the horn"}

var actions = [...]string{"lay in", "ate", "killed", "worried", "tossed", "milked", "kissed", "married", "woke", "kept", "belonged to"}

// Verse returns the ith verse of the song "This is the house that Jack built".
func Verse(verse int) string {
	var buffer bytes.Buffer

	buffer.WriteString("This is the ")
	buffer.WriteString(subjects[verse-1])

	buffer.WriteString(verseLines(verse))
	buffer.WriteString(".")

	return buffer.String()
}

func verseLines(line int) string {
	if line == 1 {
		return ""
	}

	var buffer bytes.Buffer
	buffer.WriteString("\nthat ")
	buffer.WriteString(actions[line-2])
	buffer.WriteString(" the ")
	buffer.WriteString(subjects[line-2])

	buffer.WriteString(verseLines(line - 1))

	return buffer.String()
}

// Song returns the verses of the "This is the house that Jack built" song.
func Song() string {
	var buffer bytes.Buffer

	for i := 1; i <= len(subjects); i++ {
		if i > 1 {
			buffer.WriteString("\n\n")
		}
		buffer.WriteString(Verse(i))
	}

	return buffer.String()
}
