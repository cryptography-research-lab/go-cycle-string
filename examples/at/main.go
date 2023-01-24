package main

import (
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
)

func main() {

	cycleString := cycle_string.NewCycleString("CC中文")

	// 按照字节下标获取，这两个方法是等价的
	fmt.Println("At: " + string(cycleString.At(2)))
	fmt.Println("ByteAt: " + string(cycleString.ByteAt(2)))

	// 按照字符下标获取，这两个方法是等价的
	fmt.Println("RuneAt: " + string(cycleString.RuneAt(2)))
	fmt.Println("CharAt: " + string(cycleString.CharAt(2)))

	// Output:
	// At: ä
	// ByteAt: ä
	// RuneAt: 中
	// CharAt: 中

}
