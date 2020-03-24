package problem0056

import "sort"

func merge(intervals [][]int) [][]int {
	var result [][]int
	if len(intervals) == 0 {
		return result
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastItem := result[len(result)-1]
		if lastItem[1] >= intervals[i][0] {
			start := lastItem[0]
			end := max(lastItem[1], intervals[i][1])
			result[len(result)-1] = []int{start, end}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
