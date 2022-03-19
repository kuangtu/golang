# 6 方法

Go语言从设计伊始，就不支持经典的面向对象语法元素，比如类、对象、继承等等。但保留了名为“方法（method）”的语法元素。方法能给用户定义的类型添加新的行为。

## 6.1 方法声明

在函数声明时，在其名字之前放上一个变量，即是一个方法。形式如下：

![go方法声明](jpg/go方法声明.jpg)

这个附加的receiver参数会将该函数附加到这种类型上，类似"this"，**将函数与接收者绑定。**

```go
package geometry

import "math"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

### 6.1.1 约束

方法接收器（receiver）参数、函数 / 方法参数，以及返回值变量对应的作用域范围，都是函数 / 方法体对应的显式代码块。因此，receiver 部分的参数名不能与方法参数列表中的形参名，以及具名返回值中的变量名存在冲突，必须在这个方法的作用域中具有唯一性。

- receiver 参数的基类型本身不能为指针类型或接口类型；

- 方法声明要与 receiver 参数的基类型声明放在同一个包内；
- 我们不能为原生类型（诸如 int、float64、map 等）添加方法；
- 不能跨越 Go 包为其他包的类型声明新方法。

### 6.1.2 调用

```go
package main

import (
    "fmt"
    _ "io"
)

type T struct{}

func (t T) M(n int) {
    fmt.Println(n)
}

func main() {
    var t T
    t.M(1) // 通过类型T的变量实例调用方法M

    p := &T{}
    p.M(2) // 通过类型*T的变量实例调用方法M
}
```

### 6.1.3 参数类型

receiver有两种类型的接收者：值接收者和指针接收者。（类似C语言的传值、传地址）

- 如果使用值接收者，调用时会使用这个值的一个副本来执行。
- 如果使用指针接收者，调用时使用实际值来调用方法。

假设有方法M1和M2。

```go
func (t T) M1() <=> F1(t T)
func (t *T) M2() <=> F2(t *T)
```

M1 方法是 receiver 参数类型为 T 的一类方法的代表，而 M2 方法则代表了 receiver 参数类型为 *T 的另一类。

#### 6.1.3.1 值拷贝传递

以 T 作为 receiver 参数类型时，M1 方法等价转换为F1(t T)。我们在 F1 函数的实现中对参数 t 做任何修改，都只会影响副本，而不会影响到原 T 类型实例。

#### 6.1.3.2 指针传递

以 *T 作为 receiver 参数类型时，M2 方法等价转换为F2(t *T)。传递给 F2 函数的 t 是 T 类型实例的地址，这样 F2 函数体中对参数 t 做的任何修改，都会反映到原 T 类型实例上。

#### 6.1.3.3 参数选择原则

- 如果 Go 方法要把对 receiver 参数代表的类型实例的修改，反映到原类型实例上，那么我们应该选择 *T 作为 receiver 参数的类型。

无论是 T 类型实例，还是 *T 类型实例，都既可以调用 receiver 为 T 类型的方法，也可以调用 receiver 为 *T 类型的方法。

