package memtab

import "encoding/binary"

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
	//这个8是seq，其组成为seq<<8|value
	internalKeySize := len(key) + 8
	valueSize := len(value)
	//key的格式为：编码格式为 internal_key_size| key_data(InternalKey 格式) | value_size | value_data
	encodedLen := 4 + internalKeySize + valueSize + 4
	buf := make([]byte, encodedLen)
	offset := 0
	binary.LittleEndian.PutUint32(buf[offset:], uint32(internalKeySize))
	offset += 4
	copy(buf[offset:], key)
	offset += len(key)
	//这里取值，valuetype只有可能为1和2
	binary.LittleEndian.PutUint64(buf[offset:], (uint64(seq)<<8)|uint64(valuetype))
	offset += 8
	binary.LittleEndian.PutUint32(buf[offset:], uint32(valueSize))
	offset += 4
	copy(buf[offset:], value)
	return &InternalKey{
		rep: buf,
	}

}
