package problem0169

func majorityElement(nums []int) int {
	array := make(map[int]int, len(nums))
	for _, v := range nums {
		array[v]++
		if array[v] > len(nums)/2 {
			return v
		}
	}

	return -1
}

// boyerMooreVote
// Time: O(n)
// Space: O(1)
func boyerMooreVote(nums []int) int {
	major, count := 0, 0

	for _, v := range nums {
		if count == 0 {
			major = v
		}

		if major == v {
			count++
		} else {
			count--
		}
	}

	return major
}
