package zigzagconversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := make([]byte, len(s))
	lastX, lastY := toCoordinates(len(s)-1, numRows)

	for i := range s {
		p := permute(i, numRows, lastX, lastY)
		result[p] = s[i]
	}

	return string(result)
}

func permute(n, numRows, lastX, lastY int) int {
	x, y := toCoordinates(n, numRows)

	blockSize := (numRows - 1)
	if x == 0 {
		return y / blockSize
	}

	var result int

	nFullBlocks := (lastY + 1) / blockSize
	result += (1 + 2*(x-1)) * nFullBlocks

	k := 2
	if x == numRows-1 {
		k = 1
	}
	result += k * (y / blockSize)
	if y%blockSize != 0 {
		result++
	}

	extraColumns := (lastY + 1) % (numRows - 1)
	if extraColumns == 0 {
		return result
	}

	if extraColumns == 1 && lastX < x {
		result += lastX + 1
	} else {
		result += x
	}

	if extraColumns > numRows-x {
		result += extraColumns - (numRows - x)
	}

	return result
}

func toCoordinates(n, numRows int) (x, y int) {
	n++
	y = n / (2*numRows - 2) * (numRows - 1)

	r := n % (2*numRows - 2)
	if r > 0 {
		y++
		r -= numRows
	}
	if r > 0 {
		y += r
	}
	y--

	x = (n - 1) % (2*numRows - 2)
	if x >= numRows {
		x = numRows - 1 - (x+1)%numRows
	}

	return x, y
}
