package utils

//相当于一个接口
type Comparator func(a, b interface{}) int

func IntComparator(a, b interface{}) int {
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
