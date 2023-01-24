package cycle_string

import (
	"encoding/json"
	"github.com/golang-infrastructure/go-iterator"
	"strings"
)

// CycleString 循环字符串
type CycleString struct {

	// 原始的字符串，就是对这个字符串进行周期重复
	baseString string

	// 原始字符串的字符表示形式，用于一些中文场景的处理
	baseStringRuneSlice []rune
}

var _ json.Marshaler = &CycleString{}
var _ json.Unmarshaler = &CycleString{}
var _ iterator.Iterable[rune] = &CycleString{}

// NewCycleString 创建一个循环字符串
func NewCycleString(baseString string) *CycleString {
	return &CycleString{
		baseString:          baseString,
		baseStringRuneSlice: []rune(baseString),
	}
}

// ------------------------------------------------- --------------------------------------------------------------------

// RealRuneLength 返回真实的字符长度
func (x *CycleString) RealRuneLength() int {
	return len(x.baseStringRuneSlice)
}

// RealByteLength 返回真实的字节长度
func (x *CycleString) RealByteLength() int {
	return len(x.baseString)
}

// ------------------------------------------------- --------------------------------------------------------------------

// At 返回给定位置的字节，同ByteAt
func (x *CycleString) At(index int) byte {
	return x.ByteAt(index)
}

// ByteAt 获取给定下标的字节，会按照字节长度获取
func (x *CycleString) ByteAt(index int) byte {
	if index < 0 {
		panic("Out of Index")
	}
	targetIndex := index % len(x.baseString)
	return x.baseString[targetIndex]
}

// RuneAt 获取给定下标的字符，会按照字符长度计算下标
func (x *CycleString) RuneAt(index int) rune {
	if index < 0 {
		panic("Out of Index")
	}
	targetIndex := index % len(x.baseStringRuneSlice)
	return x.baseStringRuneSlice[targetIndex]
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

// Iterator 返回一个当前对象的迭代器，当然这个迭代器是没有尽头的
func (x *CycleString) Iterator() iterator.Iterator[rune] {
	return NewCycleStringIterator(x)
}

// ------------------------------------------------- --------------------------------------------------------------------

// 转为string的时候只返回原始的字符串
func (x *CycleString) String() string {
	return x.baseString
}

// ------------------------------------------------- --------------------------------------------------------------------

// CycleStringBaseStringJsonFieldName JSON序列化后存储真实字符串的字段名称，用于让序列化后的JSON能够具有一定的辨识和阅读性
const CycleStringBaseStringJsonFieldName = "CycleStringBaseString"

// MarshalJSON JSON序列化
func (x *CycleString) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m[CycleStringBaseStringJsonFieldName] = x.baseString
	return json.Marshal(m)
}

// UnmarshalJSON JSON反序列化
func (x *CycleString) UnmarshalJSON(bytes []byte) error {
	m := make(map[string]string)
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}
	x.baseString = m[CycleStringBaseStringJsonFieldName]
	x.baseStringRuneSlice = []rune(x.baseString)
	return nil
}

// ------------------------------------------------- --------------------------------------------------------------------
