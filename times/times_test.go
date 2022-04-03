package times

import (
	"testing"
	"time"
)

func Test_PlusMonths(test *testing.T) {

	date := time.Date(2022, 3, 31, 0, 0, 0, 0, time.UTC)

	actual := PlusMonths(date, 1)

	expected := time.Date(2022, 4, 30, 0, 0, 0, 0, time.UTC)

	if actual != expected {
		test.Errorf("expected %v but was %v", expected, actual)
	}
}

func Test_MinusMonths(test *testing.T) {

	date := time.Date(2022, 3, 31, 0, 0, 0, 0, time.UTC)

	actual := MinusMonths(date, 1)

	expected := time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC)

	if actual != expected {
		test.Errorf("expected %v but was %v", expected, actual)
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
