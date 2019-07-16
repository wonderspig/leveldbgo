package memtab

type ValueType int

const (
	TypeDelete ValueType = 1
	TypeValue  ValueType = 2
)

type InternalKey struct {
	rep []byte
}

func newInternalKey(seq SequenceNumber, valuetype ValueType, key, value []byte) *InternalKey {
	//因为还有seq的长度，seq为uint64位的整数，因此我们这里需要加上他，valuetype为0或1，因此这里可以
	//加到8位的后面，即seq<<8|value
	internalKeySize := len(key) + 8
	valueSize := len(value)
	//这个是编码的长度
	encodedLen := 4 + internalKeySize + 4 + valueSize
	buf := make([]byte, encodedLen)

}
