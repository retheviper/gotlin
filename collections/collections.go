package collections

import "reflect"

// IndexOf Returns the index of the first occurrence of the specified element in the slice, or -1 if the specified element is not contained in the slice.
func IndexOf[T any](slice []*T, element *T) int {
	for index, value := range slice {
		if reflect.DeepEqual(value, element) {
			return index
		}
	}
	return -1
}

// Map Returns a slice containing the results of applying the given transform function to each element in the original slice.
func Map[T any, R any](slice []*T, transform func(*T) *R) []*R {
	result := make([]*R, len(slice))
	for index, value := range slice {
		result[index] = transform(value)
	}
	return result
}

// MapIndexed Returns a slice containing the results of applying the given transform function to each element and its index in the original slice.
func MapIndexed[T any, R any](slice []*T, transform func(int, *T) *R) []*R {
	result := make([]*R, len(slice))
	for index, value := range slice {
		result[index] = transform(index, value)
	}
	return result
}

// Partition Splits the original collection into a pair of slices, where first slice contains elements for which predicate yielded true, while second slice contains elements for which predicate yielded false.
func Partition[T any](slice []*T, predicate func(*T) bool) ([]*T, []*T) {
	matches := make([]*T, 0)
	unmatched := make([]*T, 0)
	for _, value := range slice {
		if predicate(value) {
			matches = append(matches, value)
		} else {
			unmatched = append(unmatched, value)
		}
	}
	return matches, unmatched
}

// Plus Returns an array containing all elements of the original slice and then all elements of the given elements slice.
func Plus[T any](slice []*T, other []*T) []*T {
	size := len(slice)
	result := make([]*T, size+len(other))
	for index, value := range slice {
		result[index] = value
	}
	for index, value := range other {
		result[size+index] = value
	}
	return result
}

// Reversed Returns a reversed view of the original slice.
func Reversed[T any](slice []*T) []*T {
	size := len(slice)
	result := make([]*T, size)
	i := 0
	for j := size - 1; j >= 0; j-- {
		result[i] = slice[j]
		i++
	}
	return result
}

// Filter Returns a slice containing only elements matching the given predicate.
func Filter[T any](slice []*T, predicate func(*T) bool) []*T {
	return filter(slice, predicate, false)
}

// FilterNot Returns a slice containing all elements not matching the given predicate.
func FilterNot[T any](slice []*T, predicate func(*T) bool) []*T {
	return filter(slice, predicate, true)
}

// FilterIndexed Returns a slice containing only elements matching the given predicate.
func FilterIndexed[T any](slice []*T, predicate func(int, *T) bool) []*T {
	result := make([]*T, 0)
	for index, value := range slice {
		if predicate(index, value) {
			result = append(result, value)
		}
	}
	return result
}

// FilterNotNil Returns a slice containing all elements that are not nil.
func FilterNotNil[T any](slice []*T) []*T {
	return filter(slice, func(s *T) bool {
		return s != nil
	}, true)
}

func filter[T any](slice []*T, predicate func(*T) bool, predicateShouldBeTrue bool) []*T {
	result := make([]*T, 0)
	if predicateShouldBeTrue {
		for _, value := range slice {
			if predicate(value) {
				result = append(result, value)
			}
		}
	} else {
		for _, value := range slice {
			if !predicate(value) {
				result = append(result, value)
			}
		}
	}
	return result
}

// Find Returns the first element matching the given predicate, or nil if no such element was found.
func Find[T any](slice []*T, predicate func(*T) bool) *T {
	for _, value := range slice {
		if predicate(value) {
			return value
		}
	}
	return nil
}

// All Returns true if all elements match the given predicate.
func All[T any](slice []*T, predicate func(*T) bool) bool {
	for _, value := range slice {
		if !predicate(value) {
			return false
		}
	}
	return true
}

// Any Returns true if slice has at least one element.
func Any[T any](slice []*T, predicate func(*T) bool) bool {
	return Find(slice, predicate) != nil
}

// None Returns true if the slice has no element.
func None[T any](slice []*T, predicate func(*T) bool) bool {
	return !All(slice, predicate)
}
