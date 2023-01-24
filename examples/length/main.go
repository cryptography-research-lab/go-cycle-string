package main

import (
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
)

func main() {

	cycleString := cycle_string.NewCycleString("CC中文")
	fmt.Println(fmt.Sprintf("真实的字节数： %d", cycleString.RealByteLength())) // Output: 真实的字节数： 8
	fmt.Println(fmt.Sprintf("真实的字符数： %d", cycleString.RealRuneLength())) // Output: 真实的字符数： 4

}
