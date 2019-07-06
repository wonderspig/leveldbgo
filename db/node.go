package skiplist


type Node struct{
	key interface{}
	next []*Node
}


func newNode(key interface{},level int)*Node{
	return &Node{
		key:key,
		next:=make([]*Node, level)
	}
}


func (node*Node)getNext(velvel int)*Node{
	return node.next[level]
} 

func (node*Node)setNext(level int ,val* Node){
	node.next[level]=val
}