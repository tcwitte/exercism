// Package twelve provides functions that returns the song and verses from The
// Twelve Days of Christmas.
package twelve

import "bytes"

const testVersion = 1

// Song returns the verses of the song "The Twelve Days of Christmas".
func Song() string {
	var buffer bytes.Buffer
	for i := 1; i < len(numerals); i++ {
		buffer.WriteString(Verse(i))
		buffer.WriteRune('\n')
	}
	return buffer.String()
}

// Verse returns the requested verse (1 to 12) of the song "The Twelve Days of
// Christmas".
func Verse(verse int) string {
	var buffer bytes.Buffer

	buffer.WriteString("On the ")
	buffer.WriteString(numerals[verse])
	buffer.WriteString(" day of Christmas my true love gave to me")

	for i := verse; i > 0; i-- {
		buffer.WriteString(", ")
		if verse > 1 && i == 1 {
			buffer.WriteString("and ")
		}
		buffer.WriteString(presents[i])
	}

	buffer.WriteRune('.')
	return buffer.String()
}

var numerals = [...]string{1: "first", 2: "second", 3: "third", 4: "fourth",
	5: "fifth", 6: "sixth", 7: "seventh", 8: "eighth", 9: "ninth", 10: "tenth",
	11: "eleventh", 12: "twelfth"}

var presents = [...]string{1: "a Partridge in a Pear Tree",
	2: "two Turtle Doves", 3: "three French Hens", 4: "four Calling Birds",
	5: "five Gold Rings", 6: "six Geese-a-Laying", 7: "seven Swans-a-Swimming",
	8: "eight Maids-a-Milking", 9: "nine Ladies Dancing",
	10: "ten Lords-a-Leaping", 11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming"}
