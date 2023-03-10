package main

import (
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
	"time"
)

func main() {

	cycleString := cycle_string.NewCycleString("CC11001100")
	iterator := cycleString.RuneIterator()

	fmt.Println(string(iterator.NextN(3)))

	for iterator.Next() {
		fmt.Println(string(iterator.Value()))
		time.Sleep(time.Millisecond * 100)
	}
	// Output:
	// CC1
	// 1
	// 0
	// 0
	// 1
	// 1
	// ...

}
