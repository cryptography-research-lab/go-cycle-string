package main

import (
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
	"time"
)

func main() {

	cycleString := cycle_string.NewCycleString("CC11001100")
	iterator := cycleString.Iterator()
	for iterator.Next() {
		fmt.Println(string(iterator.Value()))
		time.Sleep(time.Millisecond * 100)
	}
	// Output:
	// C
	// C
	// 1
	// 1
	// 0
	// 0
	// 1
	// 1
	// ...

}
