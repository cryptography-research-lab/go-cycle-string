package cycle_string

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycleString_At(t *testing.T) {
	s := "AAcc11001100"
	cycleString := NewCycleString(s)
	b := cycleString.At(len(s)*10000 + 1)
	t.Log(string(b))
	assert.Equal(t, "A", string(b))
}

func TestCycleString_ByteAt(t *testing.T) {
	s := "AAcc11001100"
	cycleString := NewCycleString(s)
	b := cycleString.At(len(s)*10000 + 1)
	t.Log(string(b))
	assert.Equal(t, "A", string(b))
}

func TestCycleString_CharAt(t *testing.T) {
	s := "CC是个中国人"
	cycleString := NewCycleString(s)
	b := cycleString.RuneAt(4)
	//t.Log(string(b))
	assert.Equal(t, "中", string(b))
}

func TestCycleString_RuneAt(t *testing.T) {
	s := "CC是个中国人"
	cycleString := NewCycleString(s)
	b := cycleString.RuneAt(4)
	//t.Log(string(b))
	assert.Equal(t, "中", string(b))
}

func TestCycleString_String(t *testing.T) {
	s := "CC是个中国人"
	cycleString := NewCycleString(s)
	b := cycleString.String()
	assert.Equal(t, s, b)
}

func TestCycleString_SubString(t *testing.T) {
	s := "CC11001100"
	cycleString := NewCycleString(s)
	b := cycleString.SubString(10, 12)
	assert.Equal(t, "CC", b)
}

func TestCycleString_SubStringRune(t *testing.T) {
	s := "CC是个中国人"
	cycleString := NewCycleString(s)
	b := cycleString.SubStringRune(10, 20)
	assert.Equal(t, "个中国人CC是个中国", b)
	//t.Log(b)
}

func TestCycleString_MarshalJSON(t *testing.T) {
	s := "CC是个中国人"
	cycleString := NewCycleString(s)
	b, err := cycleString.MarshalJSON()
	assert.Equal(t, []byte("{\"s\":\"CC是个中国人\"}"), b)
	assert.Nil(t, err)
}

func TestCycleString_UnmarshalJSON(t *testing.T) {
	s := "{\"s\":\"CC是个中国人\"}"
	cycleString := &CycleString{}
	err := json.Unmarshal([]byte(s), &cycleString)
	assert.Nil(t, err)
	assert.Equal(t, "CC是个中国人", cycleString.s)
}

func TestNewCycleString(t *testing.T) {
	s := "CC11001100"
	cycleString := NewCycleString(s)
	assert.NotNil(t, cycleString)
}
