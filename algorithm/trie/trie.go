package trie

// trie树的节点
type TrieNode struct {
	Data         byte
	Children     []*TrieNode
	IsEndingChar bool
}

type TrieTree struct {
	Root *TrieNode
}

var trieTree = &TrieTree{&TrieNode{Data: '/'}}

// 往Trie树中插入一个字符串
func (root *TrieNode) Insert(text string) {
	p := root

	for _, data := range text {
		index := int(data - 'a')
		if p.Children[index] == nil {
			newNode := &TrieNode{Data: byte(data)}
			p.Children[index] = newNode
		}
		p = p.Children[index]
	}
	p.IsEndingChar = true
}

// 查找某个节点
func (root *TrieNode) Find(text string) bool {
	p := root

	for _, data := range text {
		index := int(data - 'a')
		if p.Children[index] == nil {
			return false
		}
		p = p.Children[index]
	}

	//部分前缀匹配
	if !p.IsEndingChar {
		return false
	}
	return true
}
