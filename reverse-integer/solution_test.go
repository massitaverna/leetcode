package reverseinteger

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	var r int

	r = reverse(123)
	if r != 321 {
		t.Errorf("expected 321, got %d", r)
	}

	r = reverse(-123)
	if r != -321 {
		t.Errorf("expected -321, got %d", r)
	}
	
	r = reverse(10)
	if r != 1 {
		t.Errorf("expected 1, got %d", r)
	}

	r = reverse(120)
	if r != 21 {
		t.Errorf("expected 21, got %d", r)
	}

	r = reverse(math.MaxInt32-100)
	if r != 0 {
		t.Errorf("expected 0, got %d", r)
	}
}
