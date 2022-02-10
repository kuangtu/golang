# 泛型编程

2022年3月Go核心团队准备发布Go 1.18版本。该版本引入了 Go 语言开源以来最大的语法特性变化：泛型（generic）。



## 什么是泛型

简单讲就是将算法与类型解耦，实现算法更广泛的复用。已两个数相加为例：

```go

func Add(a, b int32) int32 {
  return a + b
}
```

上面的例子适用于int32类型的整数，如果是int64、int等类型进行加法运算。还需要实现对应的函数。

基于“泛型”思想，需要将算法和类型解耦，实现方式如下：

```go

func Add[T constraints.Integer](a, b T) T {
    return a + b
}
```

可以通过上述的泛型版函数去进行各种整型类型的加法运算。



## Go泛型设计的简史

从2009年开始了对泛型进行探索，直到2021年2月正式的提案才被接受。“曲折的道路”。



## Go泛型基本语法

Go 泛型的核心是类型参数（type parameter）。

### 类型参数（type parameter）

类型参数是在函数声明、方法声明的 receiver 部分或类型定义的类型参数列表中，声明的（非限定）类型名称。

类型参数在声明中充当了一个未知类型的占位符（placeholder），在泛型函数或泛型类型实例化时，类型参数会被一个类型实参替换。

类型参数列表：

```go

func GenericFoo[P aConstraint, Q anotherConstraint](x,y P, z Q)
```

这里，P，Q 是类型形参的名字，也就是类型，aConstraint，anotherConstraint 代表类型参数的约束（constraint），我们可以理解为对类型参数可选值的一种限定。

类型参数列表位于函数名与函数参数列表之间，通过一个方括号括起。**不支持变长类型参数。**

### 类型合适确定

那么 P、Q 的类型什么时候才能确定呢？这就要等到泛型函数具化（instantiation）时才能确定。另外，按惯例，类型参数（type parameter）的名字都是首字母大写的，通常都是用单个大写字母命名。

### 约束（constraint）

约束（constraint）规定了一个类型实参（type argument）必须满足的条件要求。在 Go 泛型中，我们使用 interface 类型来定义约束。Go 接口类型的定义也进行了扩展：

```go

type C1 interface {
    ~int | ~int32
    M1()
}

type T struct{}
func (T) M1() {
}

type T1 int
func (T1) M1() {
}

func foo[P C1](t P)() {
}

func main() {
    var t1 T1
    foo(t1)
    var t T
    foo(t) // 编译器报错：T does not implement C1
}
```

C1接口类型中声明了M1方法，同时声明了两个实参的类型：(~int | ~int32)，通过“|”分隔。

另外还定义了两个自定义类型T和T1，都实现了方法M1()，但是T的底层类型为struct，T1的底层类型为int，这样就导致了虽然 T 类型满足了约束 C1 的方法集合，但类型 T 因为底层类型并不是 int 或 int32 而不满足约束 C1，这也就会导致foo(t)调用在编译阶段报错。



### 类型具化（instantiation）

声明了泛型函数后，接下来就要调用泛型函数来实现具体的业务逻辑。现在我们就通过一个泛型版本 Sort 函数的调用例子，

```go

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {
}

type book struct{}
func (x book) Less(y book) bool {
        return true
}

func main() {
    var bookshelf []book
    Sort[book](bookshelf) // 泛型函数调用
}
```

过程如下：

- Sort[book]，发现要排序的对象类型为 book；
- 检查 book 类型是否满足约束要求（也就是是否实现了约束定义中的 Less 方法）
- 将泛型函数 Sort 具化为一个新函数，这里我们把它起名为 booksort，其函数原型为 func([]book)。本质上 booksort := Sort[book]。
- 用 booksort（bookshelf），整个过程只需要检查传入的函数实参（bookshelf）的类型与 booksort 函数原型中的形参类型（[]book）是否匹配就可以了。