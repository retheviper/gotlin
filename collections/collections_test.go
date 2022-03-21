package collections

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type Person struct {
	name string
	age  int
}

func Test_IndexOf(test *testing.T) {

	slice := testSlice()

	actual := IndexOf(slice, &Person{
		name: "Test3",
		age:  30,
	})

	expected := 2

	if actual != expected {
		test.Errorf("expected %v but was %v", expected, actual)
	}
}

func Test_Map(test *testing.T) {

	type Client struct {
		name string
		age  int
	}

	slice := testSlice()

	actual := Map(slice, func(p *Person) *Client {
		return &Client{
			name: p.name,
			age:  p.age,
		}
	})

	expected := []*Client{
		{
			name: "Test1",
			age:  10,
		}, {
			name: "Test2",
			age:  20,
		}, {
			name: "Test3",
			age:  30,
		}, {
			name: "Test4",
			age:  40,
		},
	}

	checkSliceEquals(expected, actual, test)
}

func Test_MapIndexed(test *testing.T) {

	type People struct {
		name   string
		number int
	}

	slice := testSlice()

	actual := MapIndexed(slice, func(index int, p *Person) *People {
		return &People{
			name:   p.name,
			number: index + 1,
		}
	})

	expected := []*People{
		{
			name:   "Test1",
			number: 1,
		}, {
			name:   "Test2",
			number: 2,
		}, {
			name:   "Test3",
			number: 3,
		}, {
			name:   "Test4",
			number: 4,
		},
	}

	checkSliceEquals(expected, actual, test)
}

func Test_Partition(test *testing.T) {

	type Citizen struct {
		name    string
		married bool
	}

	slice := []*Citizen{
		{
			name:    "Test1",
			married: true,
		}, {
			name:    "Test2",
			married: false,
		}, {
			name:    "Test3",
			married: true,
		}, {
			name:    "Test4",
			married: false,
		},
	}

	actualMarried, actualSingle := Partition(slice, func(p *Citizen) bool {
		return p.married
	})

	expectedMarried := []*Citizen{
		{
			name:    "Test1",
			married: true,
		}, {
			name:    "Test3",
			married: true,
		},
	}

	expectedSingle := []*Citizen{
		{
			name:    "Test2",
			married: false,
		}, {
			name:    "Test4",
			married: false,
		},
	}

	checkSliceEquals(expectedMarried, actualMarried, test)
	checkSliceEquals(expectedSingle, actualSingle, test)
}

func Test_Plus(test *testing.T) {

	slice := []*Person{
		{
			name: "Test1",
			age:  10,
		}, {
			name: "Test2",
			age:  20,
		},
	}

	other := []*Person{
		{
			name: "Test3",
			age:  30,
		}, {
			name: "Test4",
			age:  40,
		},
	}

	actual := Plus(slice, other)

	expected := testSlice()

	checkSliceEquals(expected, actual, test)
}

func Test_Reversed(test *testing.T) {

	slice := testSlice()

	actual := Reversed(slice)

	expected := []*Person{
		{
			name: "Test4",
			age:  40,
		}, {
			name: "Test3",
			age:  30,
		},
		{
			name: "Test2",
			age:  20,
		},
		{
			name: "Test1",
			age:  10,
		},
	}

	checkSliceEquals(expected, actual, test)
}

func Test_Filter(test *testing.T) {

	slice := testSlice()

	actual := Filter(slice, func(p *Person) bool {
		return p.age > 25
	})

	expected := []*Person{
		{
			name: "Test1",
			age:  10,
		},
		{
			name: "Test2",
			age:  20,
		},
	}

	checkSliceEquals(expected, actual, test)
}

func Test_FilterNot(test *testing.T) {

	slice := testSlice()

	actual := FilterNot(slice, func(p *Person) bool {
		return p.age > 25
	})

	expected := []*Person{
		{
			name: "Test3",
			age:  30,
		},
		{
			name: "Test4",
			age:  40,
		},
	}

	checkSliceEquals(expected, actual, test)
}

func Test_FilterIndexed(test *testing.T) {

	slice := testSlice()

	actual := FilterIndexed(slice, func(i int, p *Person) bool {
		return i%2 != 0
	})

	expected := []*Person{
		{
			name: "Test2",
			age:  20,
		},
		{
			name: "Test4",
			age:  40,
		},
	}

	checkSliceEquals(expected, actual, test)
}

func Test_FilterNotNil(test *testing.T) {

	slice := []*Person{
		{
			name: "Test1",
			age:  10,
		}, nil, {
			name: "Test3",
			age:  30,
		}, nil,
	}

	actual := FilterNotNil(slice)

	expected := []*Person{
		{
			name: "Test1",
			age:  10,
		},
		{
			name: "Test3",
			age:  30,
		},
	}

	checkSliceEquals(expected, actual, test)
}

func Test_Find(test *testing.T) {

	slice := testSlice()

	actual := Find(slice, func(p *Person) bool {
		return p.age == 20
	})

	expected := &Person{
		name: "Test2",
		age:  20,
	}

	if !reflect.DeepEqual(actual, expected) {
		test.Errorf("expected %v but was %v", expected, actual)
	}
}

func Test_All(test *testing.T) {

	slice := testSlice()

	actual := All(slice, func(p *Person) bool {
		return strings.Contains(p.name, "Test")
	})

	if !actual {
		test.Errorf("expected true but was %v", actual)
	}
}

func Test_Any(test *testing.T) {

	slice := testSlice()

	actual := Any(slice, func(p *Person) bool {
		return p.age == 30
	})

	if !actual {
		test.Errorf("expected true but was %v", actual)
	}
}

func Test_None(test *testing.T) {

	slice := testSlice()

	actual := None(slice, func(p *Person) bool {
		return p.age > 40
	})

	if !actual {
		test.Errorf("expected true but was %v", actual)
	}
}

func testSlice() []*Person {
	return []*Person{
		{
			name: "Test1",
			age:  10,
		}, {
			name: "Test2",
			age:  20,
		}, {
			name: "Test3",
			age:  30,
		}, {
			name: "Test4",
			age:  40,
		},
	}
}

func checkSliceEquals[T any](expected []*T, actual []*T, test *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		test.Errorf("expected %v but was %v", sliceToString(expected), sliceToString(actual))
	}
}

func sliceToString[T any](target []*T) string {
	var sb strings.Builder
	sb.WriteString("[")
	for index, value := range target {
		if index > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", value))
	}
	sb.WriteString("]")
	return sb.String()
}
