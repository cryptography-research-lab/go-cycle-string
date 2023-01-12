package cycle_string

import (
	"encoding/json"
	"strings"
)

// CycleString 循环字符串
type CycleString struct {
	// 原始的字符串，就是对这个字符串进行周期重复
	s string
	// 原始字符串的字符表示形式，用于一些中文场景的处理
	runeSlice []rune
}

var _ json.Marshaler = &CycleString{}
var _ json.Unmarshaler = &CycleString{}

// NewCycleString 创建一个循环字符串
func NewCycleString(s string) *CycleString {
	return &CycleString{
		s:         s,
		runeSlice: []rune(s),
	}
}

// ------------------------------------------------- --------------------------------------------------------------------

// At 返回给定位置的字节，同ByteAt
func (x *CycleString) At(index int) byte {
	return x.ByteAt(index)
}

// ByteAt 获取给定下标的字节，会按照字节长度获取
func (x *CycleString) ByteAt(index int) byte {
	targetIndex := index % len(x.s)
	return x.s[targetIndex]
}

// RuneAt 获取给定下标的字符，会按照字符长度计算下标
func (x *CycleString) RuneAt(index int) rune {
	targetIndex := index % len(x.runeSlice)
	return x.runeSlice[targetIndex]
}

// CharAt 返回给定位置的字符，同RuneAt
func (x *CycleString) CharAt(index int) rune {
	return x.RuneAt(index)
}

// ------------------------------------------------- --------------------------------------------------------------------

// SubString 获取当前字符串的子字符串
func (x *CycleString) SubString(from, to int) string {
	if to <= from || from < 0 || to <= 0 {
		return ""
	}
	// TODO 更高效的实现
	result := strings.Builder{}
	for from < to {
		result.WriteByte(x.ByteAt(from))
		from++
	}
	return result.String()
}

// SubStringRune 字符形式的子串
func (x *CycleString) SubStringRune(from, to int) string {
	if to <= from || from < 0 || to <= 0 {
		return ""
	}
	// TODO 更高效的实现
	result := strings.Builder{}
	for from < to {
		result.WriteRune(x.CharAt(from))
		from++
	}
	return result.String()
}

// ------------------------------------------------- --------------------------------------------------------------------

// 转为string的时候只返回原始的字符串
func (x *CycleString) String() string {
	return x.s
}

// ------------------------------------------------- --------------------------------------------------------------------

// MarshalJSON JSON序列化
func (x *CycleString) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["s"] = x.s
	return json.Marshal(m)
}

// UnmarshalJSON JSON反序列化
func (x *CycleString) UnmarshalJSON(bytes []byte) error {
	m := make(map[string]string)
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}
	x.s = m["s"]
	return nil
}
