package cycle_string

import (
	"fmt"
	"testing"
)

func TestNewCycleStringIterator(t *testing.T) {
	cycleString := NewCycleString("飞流直下三千尺")
	iterator := cycleString.RuneIterator()
	//for iterator.Next() {
	//	fmt.Println(string(iterator.Value()))
	//}
	fmt.Println(string(iterator.NextN(3)))
}
