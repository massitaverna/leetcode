package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
    n := len(s)

    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }

    var (
        i, j, dist, max int
        seen map[byte]struct{}
    )

    j = 1
    max = 1
    seen = make(map[byte]struct{})
    seen[s[0]] = struct{}{}
    
    for j < n {
        if _, exists := seen[s[j]]; exists {
            delete(seen, s[i])
            i++
            continue
        }

        seen[s[j]] = struct{}{}
        j++

        dist = j-i
        if dist > max {
            max = dist
        }
    }
    return max
}
