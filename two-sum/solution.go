package twosum

func twoSum(nums []int, target int) []int {
    lookupTable := make(map[int]int)

    for idx, num := range nums {
        complement := target - num
        if complementIdx, ok := lookupTable[complement]; ok {
            return []int{idx, complementIdx}
        }

        lookupTable[num] = idx
    }

    return nil
}
