package stringtointegeratoi

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	r := myAtoi("21474836460")
	if r != math.MaxInt32 {
		t.Errorf("expected %d, got %d", math.MaxInt32, r)
	}
}
