package heap

type Heap struct {
	nums  []int
	n     int
	count int
}

//init heap
func NewHeap(capacity int) *Heap {
	return &Heap{
		n:     capacity,
		nums:  make([]int, capacity+1),
		count: 0,
	}
}

//top-max heap -> heapify from down to up
func (heap *Heap) insert(data int) {
	if heap.count == heap.n {
		return
	}

	heap.count++
	heap.nums[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.nums[parent] < heap.nums[i] {
		heap.nums[parent], heap.nums[i] = heap.nums[i], heap.nums[parent]
		i = parent
		parent = i / 2
	}
}

//heapfify from up to down
func (heap *Heap) removeMax() {
	if heap.count == 0 {
		return
	}
	//swap max and last
	heap.nums[0], heap.nums[heap.count] = heap.nums[heap.count], heap.nums[0]
	heap.count--
	//heapify from up to down
	heapifyUpToDown(heap.nums, 0, heap.count)
}

//heapify
func heapifyUpToDown(nums []int, start, end int) {
	for i := start; i <= end/2; {
		maxIndex := i
		if nums[i] < nums[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= end && nums[maxIndex] < nums[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		i = maxIndex
	}
}

//heapify v2
func heap(nums []int, start, end int) {
	left := start*2 + 1
	right := left + 1

	if left > end {
		return
	}

	tmp := left
	if right <= end && nums[right] > nums[left] {
		tmp = right
	}

	if nums[tmp] > nums[start] {
		nums[start], nums[tmp] = nums[tmp], nums[start]
		heap(nums, tmp, end)
	}
}

//sort by ascend
func Sort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	h := (n - 1) >> 1
	for i := h; i >= 0; i-- {
		heapifyUpToDown(a, i, n-1)
	}

	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapifyUpToDown(a, 0, i-1)
	}
}
