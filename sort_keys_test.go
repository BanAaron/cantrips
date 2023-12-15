package cantrips

import (
	"reflect"
	"testing"
)

func TestSortKeys(t *testing.T) {
	scores := map[int]int{
		2: 2,
		1: 1,
		3: 3,
	}
	target := []int{1, 2, 3}
	keys := SortKeys(scores)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, keys)
	}
}

func TestSortKeysReverse(t *testing.T) {
	scores := map[int]int{
		2: 2,
		1: 1,
		3: 3,
	}
	target := []int{3, 2, 1}
	keys := SortKeys(scores, true)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, keys)
	}
}

func TestSortKeysEmptyMap(t *testing.T) {
	scores := map[int]int{}
	var target []int
	keys := SortKeys(scores)
	if !reflect.DeepEqual(keys, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, keys)
	}
}
