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

func (memtable *MemTable) Add(seq SequenceNumber, valueType ValueType, key, valye []byte) {

}
