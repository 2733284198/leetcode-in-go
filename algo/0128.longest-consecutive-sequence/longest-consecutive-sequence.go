package problem0128

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)

	for _, num := range nums {
		numMap[num] = true
	}

	longest := 0

	for k, _ := range numMap {

		if _, ok := numMap[k-1]; ok {
			continue
		}
		currentLength := 1

		for {
			if _, ok := numMap[k+1]; !ok {
				break
			}
			k++
			currentLength++
		}

		longest = max(longest, currentLength)
	}

	return longest
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
