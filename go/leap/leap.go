// Package leap provides leap year functionality.
package leap

const testVersion = 3

// IsLeapYear returns whether the given year is a leap year according to the Gregorian calendar.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
