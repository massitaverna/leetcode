package manacher

func longestPalindrome(s string) string {
	longestOdd := longestOddPalindrome(s)
	longestEven := longestEvenPalindrome(s)

	if len(longestOdd) > len(longestEven) {
		return longestOdd
	} else {
		return longestEven
	}
}

func longestOddPalindrome(s string) string {
	palindromesAt := make([]int, len(s))
	palindromesAt[0] = 1

	l := 0
	r := 1

	for i := 1; i < len(s); i++ {
		if i >= r {
			palindromesAt[i] = countPalindromesAt(s, i, 0)
			l = i - palindromesAt[i]
			r = i + palindromesAt[i]
			continue
		}

		j := l + r - i

		if palindromesAt[j] < r-i {
			palindromesAt[i] = palindromesAt[j]
		} else {
			palindromesAt[i] = countPalindromesAt(s, i, r-i)
		}

		if i + palindromesAt[i] > r {
			l = i - palindromesAt[i]
			r = i + palindromesAt[i]
		}
	}

	return getLongestPalindrome(s, palindromesAt)
}

func countPalindromesAt(s string, center, radius int) int {
	l := center - radius
	r := center + radius

	for {
		if l < 0 || r >= len(s) {
			break
		}
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}

	return r - center
}

func getLongestPalindrome(s string, palindromesAt []int) string {
	var max, maxIndex int

	for i, p := range palindromesAt {
		if p > max {
			max = p
			maxIndex = i
		}
	}
	return s[maxIndex-max+1:maxIndex+max]
}

func longestEvenPalindrome(s string) string {
	padded := make([]byte, 0, 2*len(s)+1)
	for i := range s {
		padded = append(padded, '*', s[i])
	}
	padded = append(padded, '*')

	longestEvenPadded := longestOddPalindrome(string(padded))

	longestEvenBytes := make([]byte, 0, len(longestEvenPadded)/2)
	for i := range longestEvenPadded {
		if longestEvenPadded[i] != '*' {
			longestEvenBytes = append(longestEvenBytes, longestEvenPadded[i])
		}
	}
	return string(longestEvenBytes)
}