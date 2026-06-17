package regularexpressionmatching

func isMatch(s string, p string) bool {
	pattern := extractPattern(p)

	memo := make([][]*bool, len(s)+1)
	for i := range memo {
		memo[i] = make([]*bool, len(pattern)+1)
	}

	return matches(s, pattern, 0, 0, memo)
}

func matches(s string, p []string, i, j int, memo [][]*bool) bool {
	if i == len(s) && j == len(p) {
		return true
	}
	if j == len(p) {
		return false
	}

	if memo[i][j] != nil {
		return *memo[i][j]
	}

	result := false
	if len(p[j]) == 2 && matches(s, p, i, j+1, memo) {
		result = true
	} else if i < len(s) && len(p[j]) == 2 && (s[i] == p[j][0] || p[j][0] == '.') && matches(s, p, i+1, j, memo) {
		result = true
	} else if i < len(s) && len(p[j]) == 1 && (s[i] == p[j][0] || p[j][0] == '.') && matches(s, p, i+1, j+1, memo) {
		result = true
	}

	memo[i][j] = &result
	return result
}

func extractPattern(p string) []string {
	pattern := make([]string, 0, len(p))
	for i := 0; i < len(p); {
		if i < len(p)-1 && p[i+1] == '*' {
			pattern = append(pattern, p[i:i+2])
			i += 2
		} else {
			pattern = append(pattern, p[i:i+1])
			i += 1
		}
	}
	return pattern
}
