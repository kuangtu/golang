# 第三章、基础数据类型

## 3.1 整数

Go语言的数值类型包括几种不同大小的整数、浮点数和复数。

Go语言同时提供了有符号和无符号类型的整数运算。这里有int8、int16、int32和int64四种截然不同大小的有符号整数类型，分别对应8、16、32、64bit大小的有符号整数，与此对应的是uint8、uint16、uint32和uint64四种无符号整数类型。

还有两种一般对应特定CPU平台机器字大小的有符号和无符号整数int和uint；其中int是应用最广泛的数值类型。这两种类型都有同样的大小，32或64bit，但是我们不能对此做任何的假设；因为不同的编译器即使在相同的硬件平台上可能产生不同的大小。**依赖机器**

还有一种无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。



## 3.2 浮点数

Go语言提供了两种精度的浮点数，float32和float64

用Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适。所有的这三个打印形式都可以指定打印的宽度和控制打印精度。

** 通过m.n%f 控制打印精度。



math包中除了提供大量常用的数学函数外，还提供了IEEE754浮点数标准中定义的特殊值的创建和测试：正无穷大和负无穷大，分别用于表示太大溢出的数字和除零的结果；还有NaN非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1).

## 3.3 类型别名

使用某个类型时，可以起另外的名字，然后在代码中使用新的名字。

Type MyInt int，MyInt就是int类型的名称，然后可以使用MyInt来操作int类型的数据。

```go
package main

import (
    "fmt"
)

type MyInt int

func main() {
    var a ,b c MyInt = 1, 2, 3
    c := a + b
    fmt.Println("the sum is:%d.", c)
}
```



## 3.4 字符串

一个字符串是一个不可改变的字节序列。字符串可以包含任意的数据，包括byte值0，但是通常是用来包含人类可读的文本。文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。

内置的len函数可以返回一个字符串中的字节数目（不是rune字符数目），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。

```GO
package main

import (
    "fmt"
)

func main() {
    s1 := "hello world"
    fmt.Println("the str len is:", len(s1))
}
```

## 3.5 byte类型

byte 是 uint8 的别名，在所有方面都等同于 uint8。

```golang
type byte = uint8
```

byte类型初始化：

```golang
    var b byte
    b = 10
    fmt.Println(b)
    
    a := [...]byte{0,1,2,3,4}
    fmt.Println(a)
```

打印出来的也是uint8的值：

![byte初始化输出](jpg/byte初始化输出.png)

### 3.5.1 byte数组转为string类型

将byte数组转发字符串：

```go
    a := []byte{0,1,2,3,4}
    fmt.Println(a)
    stra := BytesToString(a)
    //因为是不可显示的，因此输出为
    fmt.Println(stra)
    
    b := []byte{48,49,50,51,52}
    fmt.Println(b)
    strb := BytesToString(b)
    fmt.Println(strb)
```

byte数组中元素值为0~4，转为字符串之后输出是控制字符。初始化为48~52之后，字符串输出的是0~4.

![byte字符串输出](jpg/byte字符串输出.png)

### 3.5.2 string类型转为byte数组



## 3.4 运算符优先级

Go语言中关于算术运算、逻辑运算和比较运算的二元运算符，它们按照优先级递减的顺序排列：

```go
*      /      %      <<       >>     &       &^
+      -      |      ^
==     !=     <      <=       >      >=
&&
||
```

```go
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

