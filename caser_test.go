package cantrips

import (
	"testing"
)

var (
	sentance      = "the quick brown fox jumped over the lazy dog"
	sentanceLower = "the quick brown fox jumped over the lazy dog"
	sentanceUpper = "THE QUICK BROWN FOX JUMPED OVER THE LAZY DOG"
	sentanceTitle = "The Quick Brown Fox Jumped Over The Lazy Dog"
)

func TestToUpper(t *testing.T) {
	res := ToUpper(sentance)
	if res != sentanceUpper {
		t.Errorf("%s and %s are not equal", res, sentanceUpper)
	}
}

func TestToLower(t *testing.T) {
	res := ToLower(sentanceUpper)
	if res != sentanceLower {
		t.Errorf("%s and %s are not equal", res, sentanceUpper)
	}
}

func TestToTitle(t *testing.T) {
	res := ToTitle(sentanceLower)
	if res != sentanceTitle {
		t.Errorf("%s and %s are not equal", res, sentanceUpper)
	}
}
