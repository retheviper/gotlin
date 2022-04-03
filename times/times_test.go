package times

import (
	"testing"
	"time"
)

type MonthDay struct {
	Month int
	Day   int
}

func Test_PlusMonths(test *testing.T) {

	dates := map[*MonthDay]*MonthDay{}
	dates[&MonthDay{
		Month: 1,
		Day:   31,
	}] = &MonthDay{
		Month: 2,
		Day:   28,
	}
	dates[&MonthDay{
		Month: 3,
		Day:   31,
	}] = &MonthDay{
		Month: 4,
		Day:   30,
	}
	dates[&MonthDay{
		Month: 5,
		Day:   31,
	}] = &MonthDay{
		Month: 6,
		Day:   30,
	}
	dates[&MonthDay{
		Month: 8,
		Day:   31,
	}] = &MonthDay{
		Month: 9,
		Day:   30,
	}

	for fromMonthDay, toMonthDay := range dates {
		date := time.Date(2022, time.Month(fromMonthDay.Month), fromMonthDay.Day, 0, 0, 0, 0, time.UTC)
		actual := PlusMonths(date, 1)
		expected := time.Date(2022, time.Month(toMonthDay.Month), toMonthDay.Day, 0, 0, 0, 0, time.UTC)

		if actual != expected {
			test.Errorf("expected %v but was %v", expected, actual)
		}
	}
}

func Test_MinusMonths(test *testing.T) {
	dates := map[*MonthDay]*MonthDay{}
	dates[&MonthDay{
		Month: 3,
		Day:   31,
	}] = &MonthDay{
		Month: 2,
		Day:   28,
	}
	dates[&MonthDay{
		Month: 5,
		Day:   31,
	}] = &MonthDay{
		Month: 4,
		Day:   30,
	}
	dates[&MonthDay{
		Month: 7,
		Day:   31,
	}] = &MonthDay{
		Month: 6,
		Day:   30,
	}
	dates[&MonthDay{
		Month: 10,
		Day:   31,
	}] = &MonthDay{
		Month: 9,
		Day:   30,
	}
	dates[&MonthDay{
		Month: 12,
		Day:   31,
	}] = &MonthDay{
		Month: 11,
		Day:   30,
	}

	for fromMonthDay, toMonthDay := range dates {
		date := time.Date(2022, time.Month(fromMonthDay.Month), fromMonthDay.Day, 0, 0, 0, 0, time.UTC)
		actual := MinusMonths(date, 1)
		expected := time.Date(2022, time.Month(toMonthDay.Month), toMonthDay.Day, 0, 0, 0, 0, time.UTC)

		if actual != expected {
			test.Errorf("expected %v but was %v", expected, actual)
		}
	}
}

func Test_IsLeapYear_true(test *testing.T) {

	actual := IsLeapYear(2024)

	if actual != true {
		test.Errorf("expected False but was %v", actual)
	}
}

func Test_IsLeapYear_false(test *testing.T) {

	actual := IsLeapYear(2022)

	if actual != false {
		test.Errorf("expected False but was %v", actual)
	}
}
