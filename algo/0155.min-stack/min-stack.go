package problem0155

type MinStack struct {
	elements []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else if this.minStack[len(this.minStack)-1] < x { // 保证 elements 和 minStack 的长度一直
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	} else {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	this.elements = this.elements[:len(this.elements)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
