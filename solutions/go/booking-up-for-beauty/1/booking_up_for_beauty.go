package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time{
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout,date)
	return t
}

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	var passed int = time.Now().Compare(t)
	return passed == 1
}

func IsAfternoonAppointment(date string) bool{
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return fmt.Sprintf("You have an appointment on %s, %s %v, %v, at %v:%v.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

func AnniversaryDate() time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, "9/15/2012 00:00:00")
	return time.Date(time.Now().Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}
