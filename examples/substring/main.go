package main

import (
	"fmt"
	cycle_string "github.com/cryptography-research-lab/go-cycle-string"
)

func main() {

	cycleString := cycle_string.NewCycleString("中国")

	// 按字节取子串，因为是中文，所以下面就乱码了
	subString := cycleString.SubString(0, 1)
	fmt.Println("按字节数：" + subString) // Output: 按字节数：�

	// 按字符取子串，则能够完整的取到一个中文字符
	subString = cycleString.SubStringRune(0, 1)
	fmt.Println("按字符数：" + subString) // Output: 按字符数：中

	// 如果下标不合法的话则取到空字符串，不会panic
	s := cycleString.SubString(-1, 10)
	fmt.Println("越界： " + s) // Output: 越界：

}
