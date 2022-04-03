package times

import (
	"math"
	"time"
)

// PlusMonths Returns time with the given number of months added.
func PlusMonths(t time.Time, monthsToAdd int) time.Time {
	if monthsToAdd == 0 {
		return t
	}
	months := int(t.Month()) + monthsToAdd
	newYear := t.Year() + (months / 12)
	newMonth := months % 12
	if newMonth == 0 {
		newMonth = 12
	}
	return resolveLastDay(newYear, newMonth, t)
}

// MinusMonths Returns time with the given number of months subtracted.
func MinusMonths(t time.Time, monthsToSubtract int) time.Time {
	return PlusMonths(t, -monthsToSubtract)
}

// IsLeapYear Returns true is given year is leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// get smaller value between two days.
func getMinDay(day, another int) int {
	return int(math.Min(float64(day), float64(another)))
}

// resolve validated last day of month and returns time.
func resolveLastDay(year, month int, t time.Time) time.Time {
	var day int
	if month == 2 {
		if IsLeapYear(year) {
			day = getMinDay(t.Day(), 29)
		} else {
			day = getMinDay(t.Day(), 28)
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		day = getMinDay(t.Day(), 30)
	} else {
		day = t.Day()
	}
	return time.Date(year, time.Month(month), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
