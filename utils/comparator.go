package utils

import "strings"

//接口类，参照go语言圣经里面的sort.Sort接口
type Comparator interface {
	Compare(a, b interface{}) int
	Name() string
	FindShortestSeparator(start []byte, limit []byte)([]byte)
	FindShortSuccessor(key []byte)([]byte)
}

type IntComparator struct {
}

func (intcmp *IntComparator) Compare(a, b interface{}) int {
	firstNum := a.(int)
	secondNum := b.(int)
	switch {
	case firstNum > secondNum:
		return 1
	case firstNum < secondNum:
		return -1
	default:
		break
	}
	return 0
}

func (intcmp *IntComparator) Name() string {
	return "IntComparator"
}

func (intcmp *IntComparator) FindShortestSeparator(start []byte, limit []byte)([]byte) {
	return _
}

func (intcmp *IntComparator) FindShortSuccessor(key []byte) {

}

type BytewiseComparator struct {
}

//Compare BytewiseComparator借助string的位比较
func (bytewisecmp *BytewiseComparator) Compare(a, b interface{}) int {
	firstNum := a.(string)
	secondNum := b.(string)

	return strings.Compare(firstNum, secondNum)
}

//Name BytewiseComparator返回Name
func (bytewisecmp *BytewiseComparator) Name() string {
	return "BytewiseComparator"
}

func (bytewisecmp *BytewiseComparator) FindShortestSeparator(start []byte, limit []byte)([]byte) {
	startlen:=len(start)
	limitlen:=len(limit)
	minlen:=startlen<limitlen?startlen:limitlen
	diffIndex:=0
	for diffIndex<minlen && start[diffIndex]==limit[diffIndex]{
		diffIndex++
	}
	//// 如果一个字符串是另个一字符串的前缀，无需做截短操作，否则进入else。
	if diffIndex>=minlen{
	}else{
		diffByte:=start[diffIndex]
		if diffByte<byte(0xff) && diffByte+1 <limit[diffIndex]{
			start[diffIndex]++
			start=start[:diffIndex+1]
			if bytewisecmp.Compare(string(start),string(limit))>=0{
				panic("bytewise panic")
			}
		}
		 
	}
	return start
}

func (bytewisecmp *BytewiseComparator) FindShortSuccessor(key []byte)([]byte) {
	n:=len(key)
	for i:=0;i<n;i++{
		byt:=key[i]
		if byt!=byte(0xff){
			key[i]=byt+1
			key=key[:i+1]
			return key
		}
	}
	return key
}


//
type InternalKeyComparator struct{
	user_comparator_ Comparator
}


func NewInternalKeyCompator(cmp Comparator )(*InternalKeyComparator){
	return &InternalKeyComparator{
		user_comparator_:cmp,
	}
}


//Compare BytewiseComparator借助string的位比较
func (internalkeycmp *InternalKeyComparator) Compare(a, b interface{}) int {
	firstNum := a.(string)
	secondNum := b.(string)

	r:=internalkeycmp.user_comparator_.Compare(ExtractUserKey([]byte(firstNum)),ExtractUserKey([]byte(secondNum)))
	if r==0{
		//key相等的话，那么需要比较seq
		 
		
		anum:= strconv.ParseInt(firstNum[len(firstNum)-8:len(firstNum)], 10, 64)
		bnum:= strconv.ParseInt(secondNum[len(secondNum)-8:len(secondNum)], 10, 64)
		if anum>bnum{
			r=-1
		}else if(anum<bnum){
			r=1
		}
	}
	return r
}

//Name BytewiseComparator返回Name
func (internalkeycmp *InternalKeyComparator)  Name() string {
	return "BytewiseComparator"
}

func (internalkeycmp *InternalKeyComparator)  FindShortestSeparator(start []byte, limit []byte)([]byte) {
	startlen:=len(start)
	limitlen:=len(limit)
	minlen:=startlen<limitlen?startlen:limitlen
	diffIndex:=0
	for diffIndex<minlen && start[diffIndex]==limit[diffIndex]{
		diffIndex++
	}
	//// 如果一个字符串是另个一字符串的前缀，无需做截短操作，否则进入else。
	if diffIndex>=minlen{
	}else{
		diffByte:=start[diffIndex]
		if diffByte<byte(0xff) && diffByte+1 <limit[diffIndex]{
			start[diffIndex]++
			start=start[:diffIndex+1]
			if bytewisecmp.Compare(string(start),string(limit))>=0{
				panic("bytewise panic")
			}
		}
		 
	}
	return start
}

func (internalkeycmp *InternalKeyComparator)  FindShortSuccessor(key []byte)([]byte) {
	n:=len(key)
	for i:=0;i<n;i++{
		byt:=key[i]
		if byt!=byte(0xff){
			key[i]=byt+1
			key=key[:i+1]
			return key
		}
	}
	return key
}
