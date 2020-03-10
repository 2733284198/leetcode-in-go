package problem0155

type MinStack struct {
	elements []int
	h        []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if len(this.h) == 0 {
		this.h = append(this.h, x)
	} else if this.h[len(this.h)-1] < x {
		this.h = append(this.h, this.h[len(this.h)-1])
	} else {
		this.h = append(this.h, x)
	}
}

func (this *MinStack) Pop() {
	this.elements = this.elements[:len(this.elements)-1]
	this.h = this.h[:len(this.h)-1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.h[len(this.h)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
