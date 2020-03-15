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

			//fmt.Printf("path equal nums path => %v temp => %v result => %v \n", path, temp, result)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			path = append(path, nums[i])
			//fmt.Printf("before nums => %v used => %v result => %v  path => %v i => %v\n", nums, used, result, path, i)

			used[i] = true
			backTrack(path)
			used[i] = false
			path = path[:len(path)-1]
			//fmt.Printf("after nums => %v used => %v result => %v  path => %v i => %v\n", nums, used, result, path, i)

		}
	}

	//fmt.Print("1===========================================================================================1\n")

	backTrack([]int{})

	//fmt.Print("2===========================================================================================2\n")

	return result
}
