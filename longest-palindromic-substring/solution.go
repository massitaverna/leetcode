package longestpalindromicsubstring

func longestPalindrome(s string) string {
    var result string

    for i := range s {
        for offset := 0; offset < 2; offset++ {
            l := i
            r := i + offset

            for {
                if l < 0 || r >= len(s) {
                    break
                }
                if s[l] != s[r] {
                    break
                }
                if r-l >= len(result) {
                    result = s[l:r+1]
                }
                l--
                r++
            }
        }
    }

    return result
}