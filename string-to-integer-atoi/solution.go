package stringtointegeratoi

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")

	if s == "" {
		return 0
	}
	
	var sign int32
	switch s[0] {
	case '+':
		sign = 1
		s = s[1:]
	case '-':
		sign = -1
		s = s[1:]
	default:
		sign = 1
	}

	b := make([]byte, 0, len(s))
	for i, c := range s {
		if c == '0' && len(b) == 0 {
			continue
		}
		if c < '0' || c > '9' {
			break
		}
		b = append(b, s[i])
	}

	if len(b) == 0 {
		return 0
	}

	s = string(b)
	var result int32

	for i := range s {
		if sign > 0 && result > math.MaxInt32/10 {
			result = math.MaxInt32
			break
		} else if sign < 0 && result < math.MinInt32/10 {
			result = math.MinInt32
			break
		}
		result *= 10
		result += sign * int32(s[i] - '0')
		if sign > 0 && result < 0 {
			result = math.MaxInt32
			break
		} else if sign < 0 && result > 0 {
			result = math.MinInt32
			break
		}
	}

	return int(result)
}