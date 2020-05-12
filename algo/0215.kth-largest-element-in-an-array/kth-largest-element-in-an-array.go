package problem0215

import (
	"container/heap"
	"sort"
)

func findKthLargestWithSort(nums []int, k int) int {

	sort.Ints(nums)

	return nums[len(nums)-k]
}

func findKthLargest(nums []int, k int) int {
	temp := maxHeap(nums)
	h := &temp
	heap.Init(h)

	for i := 0; i < k-1; i++ {
		heap.Remove(h, 0)
	}

	return (*h)[0]
}

// maxHeap 实现了 heap 的接口
type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	// h[i] > h[j] 使得 h[0] == max(h...)
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}

func findKThLargeNumByQuick(nums []int, k int) int {
	quick(nums, 0, len(nums)-1)

	return nums[len(nums)-k]
}

func quick(nums []int, start, end int) {
	if start > end {
		return
	}

	pivot := nums[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if nums[i] < pivot {
			if splitIndex != i {
				nums[splitIndex], nums[i] = nums[i], nums[splitIndex]
			}
			splitIndex++
		}
	}

	nums[end], nums[splitIndex] = nums[splitIndex], pivot

	quick(nums, start, splitIndex-1)
	quick(nums, splitIndex+1, end)

}
