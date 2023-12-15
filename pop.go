package cantrips

import (
	"log"
)

// Pop removes last element from a slice.
//
// Returns the removed element.
//
// slice: pointer to a slice of any type
//
// index (optional): index of the element you wish to removed.
func Pop[T any](slice *[]T, index ...int) (popped T) {
	indexLength := len(index)
	sliceLength := len(*slice)
	// validate inputs
	if sliceLength == 0 {
		log.Fatalf("attempted to pop an empty list")
	}
	if indexLength > 1 {
		log.Fatalf("provided an invalid number of indexs. Number of indexes provided: %d. Only 1 index is allowed", indexLength)
	}

	// if an index has been specified
	if indexLength == 1 {
		i := index[0]
		if i > sliceLength-1 || i < 0 {
			log.Fatalf("index is out of bounds; slice length: %d, provided index: %d", sliceLength, i)
		}

		popped = (*slice)[i]
		*slice = append((*slice)[:i], (*slice)[i+1:]...)
		return
	}
	// else
	popped = (*slice)[sliceLength-1:][0]
	*slice = (*slice)[:sliceLength-1]
	return
}
