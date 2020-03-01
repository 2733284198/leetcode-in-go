package problem0217

func containsDuplicate(nums []int) bool {
	array := make(map[int]bool)
	for _, num := range nums {
		_, ok := array[num]
		if ok {
			return true
		}
		array[num] = true
	}
	return false
}
