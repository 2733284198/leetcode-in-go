package v2

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// LRUCache contains a hash map and a doubly linked list
type LRUCache struct {
	cap    int
	load   int
	keyMap map[int]*LRUNode

	front *LRUNode
	rear  *LRUNode
}

type LRUNode struct {
	prev  *LRUNode
	next  *LRUNode
	key   int
	value int
}

// Constructor initializes a new LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:    capacity,
		keyMap: make(map[int]*LRUNode),
	}
}

// Get a list node from the hash map
func (this *LRUCache) Get(key int) int {
	node, ok := this.keyMap[key]

	if !ok {
		return -1
	}

	// move node to front
	if this.front == node || (node.prev == nil && node.next == nil) {
		return node.value
	}

	this.moveNodeToFront(node)

	return node.value
}

func (this *LRUCache) insertInFront(node *LRUNode) {
	this.front.prev = node
	node.next = this.front
	this.front = node
}

func (this *LRUCache) moveNodeToFront(node *LRUNode) {
	if node == this.front {
		return
	}

	// node is the last in the list
	if node.next == nil {
		this.rear = node.prev
	} else {
		// skip next prev to node prev
		node.next.prev = node.prev
	}

	// skip prev next to node next
	node.prev.next = node.next
	node.next = this.front
	node.prev = nil
	this.front.prev = node
	this.front = node
}

// Put key and value in the LRUCache
func (this *LRUCache) Put(key int, value int) {
	if exist, ok := this.keyMap[key]; ok {
		exist.value = value
		this.moveNodeToFront(exist)
		return
	}

	node := &LRUNode{
		key:   key,
		value: value,
	}

	this.keyMap[key] = node

	// no need to evict, so place in front of list
	if this.load < this.cap {
		if this.front == nil && this.rear == nil {
			this.front = node
			this.rear = node
		} else {
			this.insertInFront(node)
		}

		this.load = this.load + 1
		return
	}

	// load is equal to capacity, so need to evict the LRU
	evicted := this.keyMap[this.rear.key]
	delete(this.keyMap, this.rear.key)

	// a single node is to be evicted
	if evicted.next == nil && evicted.prev == nil {
		this.front = node
		this.rear = node
		return
	}

	this.rear = this.rear.prev
	this.rear.next = nil
	this.insertInFront(node)
}
