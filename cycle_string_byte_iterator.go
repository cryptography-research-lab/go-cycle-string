package cycle_string

import "github.com/golang-infrastructure/go-iterator"

// CycleStringByteIterator 用于迭代一个循环字符串
type CycleStringByteIterator struct {
	cycleString *CycleString
	// 记录当前访问到的字节下标
	index int
}

var _ iterator.Iterator[byte] = &CycleStringByteIterator{}

// NewCycleStringByteIterator 基于周期字符串创建一个字符串迭代器
func NewCycleStringByteIterator(cycleString *CycleString) *CycleStringByteIterator {
	return &CycleStringByteIterator{
		cycleString: cycleString,
		index:       0,
	}
}

// Next 因为是一个无限循环，所以迭代器的方法永远返回true
func (x *CycleStringByteIterator) Next() bool {
	return true
}

// NextN 返回接下来的N个字节，用于分组加密时使用
func (x *CycleStringByteIterator) NextN(n int) []byte {
	runeSlice := make([]byte, 0)
	for n > 0 && x.Next() {
		runeSlice = append(runeSlice, x.Value())
		n--
	}
	return runeSlice
}

func (x *CycleStringByteIterator) Value() byte {
	result := x.cycleString.ByteAt(x.index)
	x.index = (x.index + 1) % x.cycleString.RealRuneLength()
	return result
}
