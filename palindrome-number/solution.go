package palindromenumber

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    digits := make([]uint8, 0, 10)
    for x > 0 {
        digit := uint8(x%10)
        digits = append(digits, digit)
        x /= 10
    }

    for i := 0; i < len(digits)/2; i++ {
        if digits[i] != digits[len(digits)-i-1] {
            return false
        }
    }
    return true
}