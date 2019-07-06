package skiplist



import (
	"math/rand"
	"sync"
	"../utils"
)

const (
	//跳表最大的层次，注意redis的hash中最大为32
	kMaxHeight =12
	//用于计算当前节点的随机层数
	kBranching = 4
)


type SkipList struct{
	level int
	head   *Node
	comparator utils.Comparator
	mu sync.RWMutex
}


func New(compare utils.Comparator) *SkipList{
	return &SkipList{
		level:0,
		head:make([]*Node,kMaxHeight),
		comparator:compare,
	}
}


func (skiplist*SkipList)Insert(key interface{}){
	skiplist.mu.Lock()
	defer skiplist.UnLock()


}


func (skiplist* SkipList)randomHeight()int{
	var height int = 1
	for height<kMaxHeight && ((rand.Int()%kBranching) == 0){
		height+=1
	}
	return height
}

func (skiplist*SkipList)findGreaterOrEqual(key interface{})*Node{

}


func (skiplist*SkipList)getMaxHeight() int {
    return level
}



func (skiplist *SkipList)equalCompare(firstKey ,sencondKey interface{})int{
	return skiplist.comparator(firstKey,sencondKey)==0
}


func (skiplist*SkipList)keyIsAfterNode(key interface{}, n* Node)bool{

}

func (skiplist*SkipList)findLessThan(key interface{}) *Node{

}

 
func (skiplist*SkipList)findLast()*Node{

}












 
  

  // Return the earliest node that comes at or after key.
  // Return NULL if there is no such node.
  //
  // If prev is non-NULL, fills prev[level] with pointer to previous
  // node at "level" for every level in [0..max_height_-1].
  Node* FindGreaterOrEqual(const Key& key, Node** prev) const;

  // Return the latest node with a key < key.
  // Return head_ if there is no such node.
  Node* FindLessThan(const Key& key) const;

  // Return the last node in the list.
  // Return head_ if list is empty.
  Node* FindLast() const;