package main

import (
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
)

func main() {

	cycleString := cycle_string.NewCycleString("CC11001100")
	fmt.Println(cycleString.String())

}
