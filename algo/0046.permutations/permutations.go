package problem0046

func permutations(nums []int) [][]int {
	var (
		result    [][]int
		used      = make([]bool, len(nums))
		backTrack func(path []int)
	)

	backTrack = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			path = append(path, nums[i])

			used[i] = true
			backTrack(path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack([]int{})

	return result
}
