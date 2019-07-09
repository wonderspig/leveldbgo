package skiplist

type Iterator struct {
	sklist *SkipList
	node   *Node
}

func newIterator(skp *SkipList) *Iterator {
	return &Iterator{
		sklist: skp,
		node:   nil,
	}
}

func (iterator *Iterator) Valid() bool {
	return iterator.node != nil
}

func (iterator *Iterator) Prev() {
	if iterator.Valid() {
		panic("current skiplist is invalid")
	}
	iterator.node = iterator.sklist.findLessThan(iterator.node.Key)
	if iterator.node == iterator.sklist.head {
		iterator.node = nil
	}
}

func (iterator *Iterator) Seek() {
	iterator.node, _ = iterator.sklist.findGreaterOrEqual(iterator.node.Key)
}

func (iterator *Iterator) SeekToFirst() {
	iterator.node = iterator.sklist.head.Next(0)
}

func (iterator *Iterator) Next() {
	if iterator.Valid() {
		panic("current skiplist is invalid")
	}
	iterator.node = iterator.node.Next(0)
}

func (iterator *Iterator) key() interface{} {
	if iterator.Valid() {
		panic("current skiplist is invalid")
	}
	return iterator.node.Key
}

func (iterator *Iterator) SeekToLast() {
	iterator.node = iterator.sklist.findLast()
	if iterator.node == iterator.sklist.head {
		iterator.node = nil
	}
}
