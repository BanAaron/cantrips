package cantrips

import "fmt"

// Pop removes last element from a slice.
//
// Returns the removed element.
//
// slice: pointer to a slice of any type
//
// index (optional): index of the element you wish to removed.
func Pop[T any](slice *[]T, index ...int) (popped T, err error) {
	indexLength := len(index)
	sliceLength := len(*slice)
	// validate inputs
	if sliceLength == 0 {
		err = fmt.Errorf("attempted to pop an empty list")
		return
	}
	if indexLength > 1 {
		err = fmt.Errorf("provided an invalid number of indexs. Number of indexes provided: %d. Only 1 index is allowed", indexLength)
		return
	}

	// if an index has been provided
	if indexLength == 1 {
		i := index[0]
		if i > sliceLength-1 || i < 0 {
			err = fmt.Errorf("index is out of bounds; slice length: %d, provided index: %d", sliceLength, i)
			return
		}

		popped = (*slice)[i]
		*slice = append((*slice)[:i], (*slice)[i+1:]...)
		return
	}

	popped = (*slice)[sliceLength-1:][0]
	*slice = (*slice)[:sliceLength-1]
	return
}
