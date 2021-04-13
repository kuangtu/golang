# 6 方法

方法能给用户定义的类型添加新的行为。

## 6.1 方法声明

在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

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

上面的代码里那个附加的参数p，叫做方法的接收器（receiver），类似"this"。**将函数与接收者绑定。**

在Go语言中，我们并不会像其它语言那样用this或者self作为接收器；我们可以任意的选择接收器的名字。由于接收器的名字经常会被使用到，所以保持其在方法间传递时的一致性和简短性是不错的主意。这里的建议是可以使用其类型的第一个字母，比如这里使用了Point的首字母p。

```go
p := Point{1, 2}
q := Point{4, 6}
fmt.Println(Distance(p, q)) // "5", function call
fmt.Println(p.Distance(q))  // "5", method call
```

Go语言中有两种类型的接收者：值接收者和指针接收者。（C语言的传值、传地址）

- 如果使用值接收者，调用时会使用这个值的一个副本来执行。
- 如果使用指针接收者，调用时使用实际值来调用方法。

使用值接收者声明的方法：

```go
bill := user{"Bill", "bill@email.com"}
bill.notify()
```

使用 bill 的值作为接收者进行调用，方法 notify 会接收到 bill的值的一个副本。**对于副本的修改，不会体现到原有值当中**

也可以使用指针来调用使用值接收者声明的方法：

```go
lisa := &user{"Lisa", "lisa@email.com"}
lisa.notify()
```

lisa指向user类型值的指针。实际上Go在代码背后执行的代码：

```go
(*lisa).notify()
```

指针被解引用为值， 这样就符合了值接收者的要求。notify操作的是一个副本。	



