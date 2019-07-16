package memtab

import (
	skiplist "../db"
)

type SequenceNumber uint64

type MemTable struct {
	table *skiplist.SkipList
}

func NewMemTable() *MemTable {
	return &MemTable{
		table: skiplist.NewSkipList(InternalKeyCompare),
	}
}

func (memtable *MemTable) Add(seq SequenceNumber, valueType ValueType, key, value []byte) {
	internalKey := newInternalKey(seq, valueType, key, value)
    memtable.table.Insert(internalKey)
}

func (memtable*MemTable)Get(key LookupKey)(value []byte,err error,bool){
	memkey:=key.memtable_key()

}
