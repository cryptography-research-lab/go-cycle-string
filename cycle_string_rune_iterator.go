package cycle_string

import "github.com/golang-infrastructure/go-iterator"

// CycleStringRuneIterator 用于迭代一个循环字符串
type CycleStringRuneIterator struct {
	cycleString *CycleString
	// 记录当前访问到的下标
	index int
}

var _ iterator.Iterator[rune] = &CycleStringRuneIterator{}

// NewCycleStringRuneIterator 基于周期字符串创建一个字符串迭代器
func NewCycleStringRuneIterator(cycleString *CycleString) *CycleStringRuneIterator {
	return &CycleStringRuneIterator{
		cycleString: cycleString,
		index:       0,
	}
}

// Next 因为是一个无限循环，所以迭代器的方法永远返回true
func (x *CycleStringRuneIterator) Next() bool {
	return true
}

// NextN 返回接下来的N个字符，用于分组加密时使用
func (x *CycleStringRuneIterator) NextN(n int) []rune {
	runeSlice := make([]rune, 0)
	for n > 0 && x.Next() {
		runeSlice = append(runeSlice, x.Value())
		n--
	}
	return runeSlice
}

func (x *CycleStringRuneIterator) Value() rune {
	result := x.cycleString.RuneAt(x.index)
	x.index = (x.index + 1) % x.cycleString.RealRuneLength()
	return result
}
