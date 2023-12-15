package cantrips

import (
	"testing"
)

var shortString = "my name is aaron"
var longString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

var shortTarget = "noraa si eman ym"
var longTarget = ".murobal tse di mina tillom tnuresed aiciffo iuq apluc ni tnus ,tnediorp non tatadipuc taceacco tnis ruetpecxE .rutairap allun taiguf ue erolod mullic esse tilev etatpulov ni tiredneherper ni rolod eruri etua siuD .tauqesnoc odommoc ae xe piuqila tu isin sirobal ocmallu noitaticrexe durtson siuq ,mainev minim da mine tU .auqila angam erolod te erobal tu tnudidicni ropmet domsuie od deS .tile gnicsipida rutetcesnoc ,tema tis rolod muspi meroL"

func TestReverseStringShortString(t *testing.T) {
	reversed := ReverseString(shortString)
	if reversed != shortTarget {
		t.Errorf("%s and %s are not equal", reversed, shortTarget)
	}
}

func TestReverseStringLongString(t *testing.T) {
	reversed := ReverseString(longString)
	if reversed != longTarget {
		t.Errorf("%s and %s are not equal", reversed, longString)
	}
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReverseString(shortString)
	}
}
func BenchmarkReverseStringLongString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ReverseString(longString)
	}
}
