package cantrips

import (
	"reflect"
	"testing"
)

func TestPopIntegers(t *testing.T) {
	numbers := make([]int, 10)
	for i := range numbers {
		numbers[i] = i + 1
	}

	target := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	returnTarget := numbers[len(numbers)-1:][0]

	name := Pop(&numbers)
	if name != returnTarget {
		t.Errorf("Invalid returned value: %v", name)
	}

	if !reflect.DeepEqual(numbers, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, numbers)
	}
}

func TestPopStings(t *testing.T) {
	names := []string{
		"aaron",
		"chris",
		"drew",
	}
	target := []string{"aaron", "chris"}
	returnTarget := "drew"

	name := Pop(&names)
	if name != returnTarget {
		t.Errorf("Invalid returned value: %v", name)
	}

	if !reflect.DeepEqual(names, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, names)
	}
}

func TestPopWithIndex(t *testing.T) {
	names := []string{"aaron", "chris", "drew"}
	target := []string{"aaron", "drew"}

	returnTarget := "chris"
	name := Pop(&names, 1)

	if !reflect.DeepEqual(target, names) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, names)
	}

	if name != returnTarget {
		t.Errorf("Invalid returned name: %v", name)
	}
}
