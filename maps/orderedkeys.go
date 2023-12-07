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

// KeysOrdered takes a map and returns a slice of keys in order. By default, keys
// will be in ascending order.
//
//	reverse: if `true` keys will be in descending order.
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
