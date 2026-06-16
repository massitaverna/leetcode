package zigzagconversion

import (
	"testing"
)

func TestPermute(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	lastX, lastY := toCoordinates(len(s)-1, numRows)
	for i := range s {
		p := permute(i, numRows, lastX, lastY)
		t.Logf("%d permuted to %d", i, p)
		if p < 0 {
			t.Errorf("%d permuted to %d, expected a non-negative integer", i, p)
		}
		if p >= len(s) {
			t.Errorf("%d permuted to %d, expected less than %d", i, p, len(s))
		}
	}
}
