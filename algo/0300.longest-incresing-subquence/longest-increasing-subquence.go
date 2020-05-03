package problem0300

// lengthOfLIS
// Time: O(n^2)
// Space: O(n)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp, res := make([]int, len(nums)), 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// lengthOfLISByBinarySearch
// Time: O(n * log(n))
// Space: O(n)
func lengthOfLISByBinarySearch(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	dp, maxLen := make([]int, len(nums)), 0

	for _, v := range nums {
		i := binarySearchInsertPosition(dp, maxLen, v) // 在长度 maxLen 的数组中，找到 v 的插入位置
		dp[i] = v                                      // 将当前值存储在 dp 中，dp 的下标来自 上面二分查找得到的位置

		if i == maxLen { // 如果 i 等于当前长度，说明是追加到新的数组 dp 的尾部
			maxLen++
		}
	}

	return maxLen
}

// 二分查找 插入位置
func binarySearchInsertPosition(d []int, maxLen, v int) int {
	low, high := 0, maxLen-1
	for low <= high {

		mid := low + (high-low)/2

		if v < d[mid] {
			high = mid - 1
		} else if v > d[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}

	return low
}
