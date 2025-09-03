type LRUCache struct {
	head     *Node
	tail     *Node
	data     map[int]*Node
	capacity int
}

type Node struct {
	prev *Node
	next *Node
	key  int
	val  int
}

func newNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

func (this *LRUCache) insert(node *Node) {
	this.data[node.key] = node
	nextNode := this.head.next
	this.head.next = node
	node.next = nextNode
	node.prev = this.head
	nextNode.prev = node
}

func (this *LRUCache) remove(node *Node) {
	delete(this.data, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}

func Constructor(capacity int) LRUCache {
	head := newNode(0, 0)
	tail := newNode(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		head:     head,
		tail:     tail,
		data:     make(map[int]*Node),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.data[key]; ok {
		this.remove(val)
		this.insert(val)
		return val.val
	}
	return -1
}

func (this *LRUCache) Put(key int, val int) {
	if value, ok := this.data[key]; ok {
		this.remove(value)
	}
	if len(this.data) == this.capacity {
		this.remove(this.tail.prev)
	}
	this.insert(newNode(key, val))

}

/**
* Your LRUCache object will be instantiated and called as such:
* obj := Constructor(capacity);
* param_1 := obj.Get(key);
* obj.Put(key,value);
 */
