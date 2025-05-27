// Package leap contains a function to find out if a given
// year is a leap year.
package leap

// IsLeapYear returns true if the given year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
