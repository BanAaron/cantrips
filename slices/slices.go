package cantrips

import "fmt"

func Pop[T any](slice *[]T) (T, error) {
	var popped T
	length := len(*slice)
	if length == 0 {
		err := fmt.Errorf("attempted to pop an empty list")
		return popped, err
	}

	popped = (*slice)[length-1:][0]
	*slice = (*slice)[:length-1]
	return popped, nil
}
