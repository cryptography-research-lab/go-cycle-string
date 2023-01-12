package cycle_string

import (
	"fmt"
	"testing"
)

func TestNewCycleStringIterator(t *testing.T) {
	cycleString := NewCycleString("飞流直下三千尺")
	iterator := cycleString.Iterator()
	for iterator.Next() {
		fmt.Println(string(iterator.Value()))
	}
}
