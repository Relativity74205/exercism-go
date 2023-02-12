package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	schedule, _ := time.Parse("1/02/2006 15:04:05", date)
	return schedule
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateParsed, _ := time.Parse("January 2, 2006 15:04:05", date)
	return dateParsed.Unix() < time.Now().Unix()
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateParsed, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return dateParsed.Hour() >= 12 && dateParsed.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	schedule, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %v, at %v.", schedule.Format("Monday, January 2, 2006"), schedule.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(2023, 9, 15, 0, 0, 0, 0, time.UTC)
}
