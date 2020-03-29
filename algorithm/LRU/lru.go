package LRU

//双向链表的节点
type Node struct {
	Pre   *Node
	Next  *Node
	Value int
	Key   int
}
type LRUCache struct {
	Capacity int
	HTable   map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		HTable:   make(map[int]*Node),
		Head:     &Node{},
		Tail:     &Node{},
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Pre = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if this == nil || len(this.HTable) == 0 {
		return -1
	}
	node, ok := this.HTable[key]
	//如果不存在，返回-1
	if !ok {
		return -1
	}
	//如果存在调整节点到头，再返回结果
	this.Remove(node)
	this.SetHead(node)

	return node.Value

}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.HTable[key]
	//已存在
	if ok {
		this.Remove(node)
		node.Value = value
	} else { //不存在
		if len(this.HTable) == this.Capacity {
			delete(this.HTable, this.Tail.Pre.Key)
			this.Remove(this.Tail.Pre)
		}
		node = &Node{Value: value, Key: key}
		this.HTable[key] = node
	}
	this.SetHead(node)
}

func (this *LRUCache) Remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) SetHead(node *Node) {
	node.Next = this.Head.Next
	node.Next.Pre = node
	this.Head.Next = node
	node.Pre = this.Head

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
