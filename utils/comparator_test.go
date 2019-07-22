package main

import (
	"fmt"
	"strings"
)

//Comparator interface，所有的父类在他的基础上进行扩展
type Comparator interface {
	Compare(a, b interface{}) int
	Name() string
	FindShortestSeparator(start []*string, limit []byte)
	FindShortSuccessor(key []*byte)
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
	return ("IntComparator")
}

func (intcmp *IntComparator) FindShortestSeparator(start []*string, limit []byte) {

}

func (intcmp *IntComparator) FindShortSuccessor(key []*byte) {

}

type MyStruct struct {
	cmp Comparator
	val int
}

func CompareStruct(a MyStruct, value int) int {
	return a.cmp.Compare(a.val, value)
}
func main() {
	intcmp := IntComparator{}
	result := intcmp.Compare(1, 1)
	fmt.Println(result)
	fmt.Println(intcmp.Name())
	//	mya := MyStruct{
	//	cmp: &IntComparator{},
	//	val: 1,
	//}
	//fmt.Println(CompareStruct(mya, 3))
	var str string = "aaa"
	str = "bb"

	fmt.Println(str)
	fmt.Println(strings.Compare("bb", "ba"))

}
