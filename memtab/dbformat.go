package memtab

import "encoding/binary"

//需要注意的是user_key,internal_key,memtable_key
type LookupKey struct {
	start_  []byte //起始位置
	kstart_ []byte //原先的变长internal_size，实际上kstart-4=start定长为4个字节
	end_    []byte //表示结束，这里可以不用他
}

func newLookupKey(user_key []byte, sequence SequenceNumber) *LookupKey {
	internalKeySize := len(user_key) + 8
	valueSize := len(value)
	//key的格式为：编码格式为 internal_key_size| key_data(InternalKey 格式)
	encodedLen := 4 + internalKeySize
	buf := make([]byte, encodedLen)
	offset := 0
	binary.LittleEndian.PutUint32(buf[offset:], uint32(internalKeySize))
	offset += 4
	copy(buf[offset:], user_key)
	offset += len(user_key)
	//这里取值，valuetype只有可能为1和2
	binary.LittleEndian.PutUint64(buf[offset:], (uint64(sequence)<<8)|uint64(valuetype))
	buf_ := make([]byte, internalKeySize)
	copy(buf_, buf[4:])
	return &LookupKey{
		start_: buf,
		end_:   buf_,
		end_:   nil,
	}

}

//用于memtable中查询
func (lookupkey *LookupKey) memtable_key() []byte {
	buf := make([]byte, len(lookupkey.start_))
	copy(buf, lookupkey.start_)
	return buf
}

//用于内部迭代器的遍历
func (lookupkey *LookupKey) internal_key() []byte {
	buf := make([]byte, len(lookupkey.kstart_))
	copy(buf, lookupkey.kstart_)
	return buf
}

//用户实际上的user_key
func (lookupkey *LookupKey) user_key() []byte {
	buf := make([]byte, len(lookupkey.kstart_)-8)
	copy(buf, lookupkey.kstart[:len(lookupkey.kstart_)-8])
	return buf
}
