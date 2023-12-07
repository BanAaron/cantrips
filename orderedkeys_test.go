package cantrips

import (
	"reflect"
	"testing"
)

func TestOrderedKeys(t *testing.T) {
	scores := map[int]int{
		2: 2,
		1: 1,
		3: 3,
	}

	target := []int{1, 2, 3}

	keys, err := KeysOrdered(scores)
	if err != nil {
		t.Errorf("KeysOrdered returned an error: %v", err)
	}

	if !reflect.DeepEqual(keys, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, keys)
	}
}

func TestKeysOrderedReverse(t *testing.T) {
	scores := map[int]int{
		2: 2,
		1: 1,
		3: 3,
	}

	target := []int{3, 2, 1}

	keys, err := KeysOrdered(scores, true)
	if err != nil {
		t.Errorf("KeysOrdered returned an error: %v", err)
	}

	if !reflect.DeepEqual(keys, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, keys)
	}
}
