# 循环字符串数据结构（CycleString）

# 一、这是什么？解决了什么问题？

这个库定义了一个名为`CycleString`的数据结构，其实就是基于一个有限长度的字符串进行无限次重复得到的一个长度无限的字符流，同时在这个数据结构的基础上提供了一些字符串操作方法以便将其当做一个正常的字符串使用。

在研究古典密码学的时候发现一个很常见的场景就是秘钥补齐，比如要加密的文本的字符长度是20，但是加密使用的秘钥字符长度是10，则需要将秘钥重复两次凑齐20个字符与要加密的文本对齐再进行加密运算，于是就萌生了一个想法，干脆就定义一个特殊的字符串结构，这个字符串结构可以看做是一个基础字符串进行无限次重复而得到的一个长度无限的字符流，如此在加密场景下则可以忽略秘钥的长度让逻辑更清晰，从而将精力放在更重要的核心逻辑上而不是处理这些边边角角的鸡毛蒜皮，不止古典密码加密算法，其它类似的场景需求都可复用此工具库。

# 二、安装

```bash
go get -u github.com/cryptography-research-lab/go-cycle-string
```

# 三、示例代码

## 3.1 创建一个无限循环字符串

此数据结构需要基于一个有限长度的字符串，比如下面的代码就是创建一个无限长度的字符串：

```go
cycleString := cycle_string.NewCycleString("CC11001100")
```

这个无限长度的字符串的前若干个字符是：

```text
CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100CC11001100
```

## 3.2 字符串真实长度

对于`CycleString`这种数据结构来说长度是没有意义的，因为它的长度是无穷，但是它所基于的字符串的长度是有限的，所以这里的真实长度都是指的它所基于的字符串的长度，这里的长度有两种，一种是字符长度，使用`RealRuneLength`来获取，一种是字节长度，使用`RealByteLength`来获取，下面是这两个方法的使用示例：

```go
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
```

## 3.3 字符 & 字节

`CycleString`本质上是一个字符串，所以也可以按照下标获取对应位置的字符，对于下标的类型，又分为按照字节的下标和按照字符的下标：

```go
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
```

## 3.4 子串 

`CycleString`本质上是一个字符串，所以也可以获取其子串，获取子串有两种类型，一种是按字节长度获取，对应的方法名是`SubString`，一种是按字符数来获取，对应的方法是`SubStringRune`，使用哪个方法请按照自己的需求决定：

```go
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
```

## 3.5 迭代器模式

`CycleString`有一个迭代器模式的实现`CycleStringIterator`，通过`Iterator`方法来获取：

```go
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
```

## 3.6 JSON序列化 & 反序列化 

此数据结构可以被正常的序列化和反序列化，只不过比较特殊的是因为`CycleString`这个数据结构本身来说是没有右边界的，而`JSON`序列化又需要它有一个明确的边界，所以`CycleString`重写了`JSON`序列化的`MarshalJSON`和`UnmarshalJSON`函数，在序列化的时候只保存`CycleString`所基于的字符串，下面的代码是一个`JSON`序列化的例子：

```go
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
```

# 四、TODO 

暂未发现更多需求，如您使用了这个库又有相关功能当前未提供，可提`Issues`区讨论。

