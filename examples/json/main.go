package main

import (
	"encoding/json"
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
)

// Foo CycleString作为Foo这个Struct的一个字段，Foo可以被安全的JSON序列化
type Foo struct {
	Bar *cycle_string.CycleString
}

func main() {

	foo := Foo{
		Bar: cycle_string.NewCycleString("CC11001100"),
	}
	marshal, err := json.Marshal(foo)
	if err != nil {
		fmt.Println("JSON序列化错误： " + err.Error())
		return
	}
	fmt.Println("序列化后的文本是： " + string(marshal)) // Output: 序列化后的文本是： {"Bar":{"CycleStringBaseString":"CC11001100"}}

	foo2 := &Foo{}
	err = json.Unmarshal(marshal, foo2)
	if err != nil {
		fmt.Println("JSON反序列化错误： " + err.Error())
		return
	}
	fmt.Println("反序列后的前30个字符： " + foo2.Bar.SubStringRune(0, 30)) // Output: 反序列后的前30个字符： CC11001100CC11001100CC11001100

}
