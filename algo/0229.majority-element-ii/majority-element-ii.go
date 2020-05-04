package problem0229

func majorityElement(nums []int) []int {

	result := []int{}

	if nums == nil || len(nums) == 0 {
		return result
	}

	c1, count1, c2, count2 := nums[0], 0, nums[0], 0

	for _, num := range nums {
		if c1 == num {
			count1++
			continue
		}

		if c2 == num {
			count2++
			continue
		}

		if count1 == 0 {
			c1 = num
			count1++
			continue
		}

		if count2 == 0 {
			c2 = num
			count2++
			continue
		}

		count1--
		count2--

	}

	count1 = 0
	count2 = 0

	for _, num := range nums {
		if c1 == num {
			count1++
		}

		if c2 == num {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		result = append(result, c1)
	}

	if count2 > len(nums)/3 {
		result = append(result, c2)
	}

	return result
}
