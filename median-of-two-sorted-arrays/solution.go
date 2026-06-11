package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return arrayMedian(nums2)
	}

	if len(nums2) == 0 {
		return arrayMedian(nums1)
	}

	medians := make(map[int]struct{})
	medianSearch(nums1, nums2, medians)
	medianSearch(nums2, nums1, medians)

	var sum float64
	for k := range medians {
		sum += float64(k)
	}

	if len(medians) == 1 {
		return sum
	}
	return sum / 2
}

func medianSearch(nums, others []int, medians map[int]struct{}) {
	start := 0
	end := len(nums) - 1

	for start <= end {
		var median *int
		start, end, median = medianSearchStep(nums, start, end, others)
		if median != nil {
			medians[*median] = struct{}{}
		}
	}
}

func medianSearchStep(nums []int, start, end int, others []int) (int, int, *int) {
	m := (start + end) / 2

	r := binarySearch(nums[m], others)
	l := binarySearch(nums[m]-1, others) + 1
	smaller := l
	greater := len(others) - r - 1
	occurrences := r - l + 1

	r = binarySearch(nums[m], nums)
	l = binarySearch(nums[m]-1, nums) + 1
	smaller += l
	greater += len(nums) - r - 1
	occurrences += r - l + 1

	if isMedian(smaller, greater, occurrences) {
		if greater > smaller {
			return m + 1, end, &nums[m]
		} else {
			return start, m - 1, &nums[m]
		}
	} else {
		if greater > smaller {
			return r + 1, end, nil
		} else {
			return start, l - 1, nil
		}
	}

}

func isMedian(smaller, greater, occurrences int) bool {
	var d int
	if greater > smaller {
		d = greater - smaller
	} else {
		d = smaller - greater
	}
	return d <= occurrences
}

func binarySearch(target int, nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		m := (start + end + 1) / 2
		if nums[m] <= target {
			start = m
		} else {
			end = m - 1
		}
	}

	if nums[start] > target {
		return -1
	}
	return start
}

func arrayMedian(nums []int) float64 {
	m := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[m-1]+nums[m]) / 2
	}
	return float64(nums[m])
}
