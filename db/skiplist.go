package skiplist

import (
	"fmt"
	"sync"

	"../utils"
)

const (
	//跳表最大的层次，注意redis的hash中最大为32
	kMaxHeight = 12
	//用于计算当前节点的随机层数
	kBranching = 4
)

type SkipList struct {
	maxHeight int
	head      *Node

	rnd        *utils.Random
	comparator utils.Comparator
	mu         sync.RWMutex
}

func NewSkipList(compare utils.Comparator) *SkipList {
	return &SkipList{
		maxHeight:  1,
		head:       newNode(nil, kMaxHeight),
		comparator: &utils.IntComparator{},
		rnd:        utils.NewRandom(3735928559),
	}
}

func (skiplist *SkipList) Insert(key interface{}) {
	skiplist.mu.Lock()
	defer skiplist.mu.Unlock()
	x, prev := skiplist.findGreaterOrEqual(key)

	if x == nil || !skiplist.equalCompare(key, x.Key) {
		fmt.Println("data structure does not allow duplicate insertion")
		return
	}
	//个人感觉这里应该先计算随机高度，然后
	height := skiplist.randomHeight()
	if height > skiplist.getMaxHeight() {
		for i := skiplist.getMaxHeight(); i < height; i++ {
			prev[i] = skiplist.head
		}
		skiplist.maxHeight = height
	}

	x = newNode(key, height)
	for i := 0; i < height; i++ {
		x.SetNext(i, prev[i].GetNext(i))
		prev[i].SetNext(i, x)
	}
}

func (skiplist *SkipList) Contain(key interface{}) bool {
	x, _ := skiplist.findGreaterOrEqual(key)
	if x != nil && skiplist.Equal(key, x.Key) {
		return true
	}
	return false
}

func (skiplist *SkipList) Equal(keya, keyb interface{}) bool {
	return (skiplist.comparator(keya, keyb) == 0)
}

func (skiplist *SkipList) randomHeight() int {
	var height int = 1
	for height < kMaxHeight && ((skiplist.rnd.Next() % kBranching) == 0) {
		height += 1
	}
	return height
}

func (skiplist *SkipList) findGreaterOrEqual(key interface{}) (*Node, []*Node) {
	//skiplist是排好序的，可以理解为插入排序
	prev := make([]*Node, kMaxHeight)
	x := skiplist.head
	level := skiplist.getMaxHeight() - 1
	for {
		next := x.Next(level)
		if skiplist.keyIsAfterNode(key, next) {
			x = next
		} else {
			if prev != nil {
				prev[level] = x
			}
			if level == 0 {
				return next, prev
			} else {
				level--
			}
		}

	}

}

func (skiplist *SkipList) getMaxHeight() int {
	return skiplist.maxHeight
}

func (skiplist *SkipList) equalCompare(firstKey, sencondKey interface{}) bool {
	return skiplist.comparator(firstKey, sencondKey) == 0
}

func (skiplist *SkipList) keyIsAfterNode(key interface{}, n *Node) bool {
	return (n != nil) && (skiplist.comparator(n.Key, key) < 0)
}

func (skiplist *SkipList) findLast() *Node {
	x := skiplist.head
	level := skiplist.getMaxHeight() - 1
	for {
		next := x.Next(level)
		if next == nil {
			if level == 0 {
				return x
			} else {
				level -= 1
			}
		} else {
			x = next
		}
	}
}

func (skiplist *SkipList) findLessThan(key interface{}) *Node {
	x := skiplist.head
	level := skiplist.getMaxHeight() - 1
	for {
		if x == skiplist.head || skiplist.comparator.Compare(x.Key, key) < 0 {

		} else {
			panic("It has a problem!The program exit")

		}
		next := x.Next(level)
		if next == nil || skiplist.comparator.Compare(x.Key, key) >= 0 {
			if level == 0 {
				return x
			} else {
				level--
			}

		} else {
			x = next
		}
	}
}
