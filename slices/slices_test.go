package cantrips

import (
	"reflect"
	"testing"
)

func TestPopSliceOfIntegers(t *testing.T) {
	numbers := make([]int, 10)
	for i := range numbers {
		numbers[i] = i + 1
	}

	target := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	returnTarget := numbers[len(numbers)-1:][0]

	v, err := Pop(&numbers)
	if err != nil {
		t.Errorf("Pop returned an error: %v", err)
	}

	if v != returnTarget {
		t.Errorf("Invalid returned value: %v", v)
	}

	if !reflect.DeepEqual(numbers, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, numbers)
	}
}
