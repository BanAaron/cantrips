package cantrips

import (
	"cmp"
	"fmt"
	"sort"
)

// ordered is an interface that says "whatever datatype implements this interface"
// must be of the type cmp.Ordered.
type ordered interface {
	cmp.Ordered
}

// KeysOrdered returns a slice of keys in ascending order.
//
// input: a map with int, float, or string keys
//
// reverse: keys will be in descending order if set to true
func KeysOrdered[Key ordered, Value any](input map[Key]Value, reverse ...bool) (keys []Key, err error) {
	reverseLength := len(reverse)
	if reverseLength > 1 {
		err = fmt.Errorf("Too many optional arguments provided. Number of option arguments provided: %d\n", reverseLength)
		return
	}
	doReverse := false

	if reverseLength > 0 {
		doReverse = reverse[0]
	}

	for k := range input {
		keys = append(keys, k)
	}

	if doReverse {
		sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })
	} else {
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	}
	return
}
