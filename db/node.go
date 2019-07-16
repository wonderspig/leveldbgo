package skiplist

type Node struct {
	Key     interface{}
	forward []*Node
}

func newNode(key interface{}, level int) *Node {
	return &Node{
		Key:     key,
		forward: make([]*Node, level),
	}
}

func (node *Node) Next(n int) *Node {
	if n < 0 {
		panic("out of bound")
	}
	return node.forward[n]
}

func (node *Node) GetNext(level int) *Node {
	return node.forward[level]
}

func (node *Node) SetNext(level int, val *Node) {
	node.forward[level] = val
}
