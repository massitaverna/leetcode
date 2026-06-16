package reverseinteger

import (
	"math"
)

func reverse(x int) int {
	var (
		n        int32 = int32(x)
		sign     int32 = 1
		result   int32
		overflow bool
	)

	if n < 0 {
		sign = -1
		n = -n
	}

	for n >= 10 {
		result += n % 10
		if result < 0 {
			overflow = true
			break
		}

		if result > math.MaxInt32/10 {
			overflow = true
			break
		}
		result *= 10

		n = n / 10
	}

	result += n % 10
	if result < 0 {
		overflow = true
	}

	if overflow {
		return 0
	}

	return int(sign * result)
}
