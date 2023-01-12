package cycle_string

import "github.com/golang-infrastructure/go-iterator"

// CycleStringIterator 用于迭代一个循环字符串
type CycleStringIterator struct {
	cycleString *CycleString
	// 记录当前访问到的下标
	index int
}

var _ iterator.Iterator[rune] = &CycleStringIterator{}

// NewCycleStringIterator 基于周期字符串创建一个字符串迭代器
func NewCycleStringIterator(cycleString *CycleString) *CycleStringIterator {
	return &CycleStringIterator{
		cycleString: cycleString,
		index:       0,
	}
}

// Next 因为是一个无限循环，所以迭代器的方法永远返回true
func (x *CycleStringIterator) Next() bool {
	return true
}

func (x *CycleStringIterator) Value() rune {
	result := x.cycleString.RuneAt(x.index)
	x.index = (x.index + 1) % x.cycleString.RealRuneLength()
	return result
}
