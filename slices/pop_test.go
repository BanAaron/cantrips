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

func TestPopStings(t *testing.T) {
	names := []string{
		"aaron",
		"chris",
		"drew",
	}
	target := []string{"aaron", "chris"}
	returnTarget := "drew"

	v, err := Pop(&names)
	if err != nil {
		t.Errorf("Pop returned an error: %v", err)
	}

	if v != returnTarget {
		t.Errorf("Invalid returned value: %v", v)
	}

	if !reflect.DeepEqual(names, target) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, names)
	}
}

func TestPopWithIndex(t *testing.T) {
	names := []string{"aaron", "chris", "drew"}
	target := []string{"aaron", "drew"}
	returnTarget := "chris"

	name, err := Pop(&names, 1)
	if err != nil {
		t.Errorf("Pop returned an error: %v", err)
	}

	if !reflect.DeepEqual(target, names) {
		t.Errorf("Resulting slice is invalid.\nExpecteed: %v\nResult: %v", target, names)
	}

	if name != returnTarget {
		t.Errorf("Invalid returned name: %v", name)
	}
}

func TestPopNegativeIndex(t *testing.T) {
	names := []string{"aaron", "chris", "drew"}

	_, err := Pop(&names, -1)
	if err == nil {
		t.Errorf("Pop did not return an error when passing in a nagative index")
	}
}
