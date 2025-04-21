package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	schedule, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return time.Now().After(schedule)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	schedule, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	if hour := schedule.Hour(); hour >= 12 && hour <= 17 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	schedule, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	formattedDate := schedule.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formattedDate)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
