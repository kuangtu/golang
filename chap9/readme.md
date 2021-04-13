# 第九章 基于共享变量的开发



## 9.1 竞争条件

​	在一个协程的程序中，程序的执行顺序只由程序的逻辑来决定。在多个协程中，也是按照既定的顺序去执行。但是通常情况下不知道分别位于两个协程中事件x和事件y的执行顺序。

竞争条件指的是程序在多个goroutine交叉执行操作时，没有给出正确的结果。

基于银行账户修改，来说明情况：

```go
// Package bank implements a bank with only one account.
package bank
var balance int
func Deposit(amount int) { balance = balance + amount }
func Balance() int { return balance }
```

如果发运行：

```go
// Alice:
go func() {
    bank.Deposit(200)                // A1
    fmt.Println("=", bank.Balance()) // A2
}()

// Bob:
go bank.Deposit(100)                 // B
```

Alice存了200，然后检查她的余额，同时Bob存了200，然后检查她的余额，同时Bob存了100。因为A1和A2是和B并发执行的，我们没法预测他们发生的先后顺序。通常认为出现三种可能性：“Alice先”，“Bob先”以及“Alice/Bob/Alice”交错执行。但是**会出现另外一种情况“，这种情况下Bob的存款会在Alice存款操作中间，在余额被读到（balance + amount）之后，在余额被更新之前（balance = ...），这样会导致Bob的交易丢失。

```go
Data race
0
A1r      0     ... = balance + amount
B      100
A1w    200     balance = ...
A2  "= 200"
```

因为balance赋值不是原始的。

**无论任何时候，只要有两个goroutine并发访问同一变量，且至少其中的一个是写操作的时候就会发生数据竞争。**



## 9.2 sync.Mutex互斥锁

sync包里的Mutex，它的Lock方法能够获取到token(这里叫锁)，并且Unlock方法会释放这个token：

```go
import "sync"

var (
    mu      sync.Mutex // guards balance
    balance int
)

func Deposit(amount int) {
    mu.Lock()
    balance = balance + amount
    mu.Unlock()
}

func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}
```



每次一个goroutine访问bank变量时（这里只有balance余额变量），它都会调用mutex的Lock方法来获取一个互斥锁。如果其它的goroutine已经获得了这个锁的话，这个操作会被阻塞直到其它goroutine调用了Unlock使该锁变回可用状态。mutex会保护共享变量。

**如果函数中分支较多，释放锁逻辑处理比较繁琐**，通过defer语法处理，我们用defer来调用Unlock，临界区会隐式地延伸到函数作用域的最后，这样我们就从“总要记得在函数返回之后或者发生错误返回时要记得调用一次Unlock”这种状态中获得了解放。

```go
func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}
```

### 9.2.1 Snyc.Once

Once类型的对象，只执行一次动作。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

```

Once类型执行Do方法：

```go
func (o *Once) Do(f func())
```

只执行一次.

## 9.3 读写锁

由于Balance函数只需要读取变量的状态，所以我们同时让多个Balance调用并发运行事实上是安全的，只要在运行的时候没有存款或者取款操作就行。在这种场景下我们需要一种特殊类型的锁，其允许多个只读操作并行执行，但写操作会完全互斥。这种锁叫作“多读单写”锁（multiple readers, single writer lock），Go语言提供的这样的锁是sync.RWMutex：

```go
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // readers lock
    defer mu.RUnlock()
    return balance
}
```



## 9.4 内存同步

在现代计算机中可能会有一堆处理器，每一个都会有其本地缓存（local cache）。为了效率，对内存的写入一般会在每一个处理器中缓冲，并在必要时一起flush到主存。这种情况下这些数据可能会以与当初goroutine写入顺序不同的顺序被提交到主存。

考虑如下代码可能的输出：

```go
var x, y int
go func() {
    x = 1 // A1
    fmt.Print("y:", y, " ") // A2
}()
go func() {
    y = 1                   // B1
    fmt.Print("x:", x, " ") // B2
}()
```

因为两个goroutine是并发执行，并且访问共享变量时也没有互斥，会有数据竞争，所以程序的运行结果没法预测的话也请不要惊讶。我们可能希望它能够打印出下面这四种结果中的一种，相当于几种不同的交错执行时的情况：

```go
y:0 x:1
x:0 y:1
x:1 y:1
y:1 x:1
```

第四行可以被解释为执行顺序A1,B1,A2,B2或者B1,A1,A2,B2的执行结果。然而实际运行时还是有些情况让我们有点惊讶：

```go
x:0 y:0
y:0 x:0
```

根据所使用的编译器，CPU，或者其它很多影响因子，这两种情况也是有可能发生的。**如何解释？**

但看看上面的例子，这已经不是现代的编译器和cpu的工作方式了。因为赋值和打印指向不同的变量，编译器可能会断定两条语句的顺序不会影响执行结果，并且会交换两个语句的执行顺序。如果两个goroutine在不同的CPU上执行，每一个核心有自己的缓存，这样一个goroutine的写入对于其它goroutine的Print，在主存同步之前就是不可见的了。

所有并发的问题都可以用一致的、简单的既定的模式来规避。所以可能的话，将变量限定在goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问。

## 9.6 竞争条件检测

Go的runtime和工具链为我们装备了一个复杂但好用的动态分析工具，竞争检查器（the race detector）。

只要在go build，go run或者go test命令后面加上-race的flag，就会使编译器创建一个你的应用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的test，并且会记录下每一个读或者写共享变量的goroutine的身份信息。

竞争检查器会检查这些事件，会寻找在哪一个goroutine中出现了这样的case，例如其读或者写了一个共享变量，这个共享变量是被另一个goroutine在没有进行干预同步操作便直接写入的。这种情况也就表明了是对一个共享变量的并发访问，即数据竞争。



## 9.8 Goroutines和线程



### 9.8.1 动态栈

每一个OS线程都有一个固定大小的内存块（一般会是2MB）来做栈，这个栈会用来存储当前正在被调用或挂起（指在调用其它函数时）的函数的内部变量。

一个goroutine的栈，和操作系统线程一样，会保存其活跃或挂起的函数调用的本地变量，但是和OS线程不太一样的是，一个goroutine的栈大小并不是固定的；栈的大小会根据需要动态地伸缩。而goroutine的栈的最大值有1GB，比传统的固定大小的线程栈要大得多，尽管一般情况下，大多goroutine都不需要这么大的栈。



### 9.8.2 Goroutine调度

OS线程会被操作系统内核调度。每几毫秒，一个硬件计时器会中断处理器，这会调用一个叫作scheduler的内核函数。这个函数会挂起当前执行的线程并将它的寄存器内容保存到内存中，检查线程列表并决定下一次哪个线程可以被运行，并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。线程切换。

和操作系统的线程调度不同的是，Go调度器并不是用一个硬件定时器，而是被Go语言“建筑”本身进行调度的。例如当一个goroutine调用了time.Sleep，或者被channel调用或者mutex操作阻塞时，调度器会使其进入休眠并开始执行另一个goroutine，直到时机到了再去唤醒第一个goroutine。因为这种调度方式不需要进入内核的上下文，所以重新调度一个goroutine比调度一个线程代价要低得多。



### 9.8.4 Goroutine没有ID号

在大多数支持多线程的操作系统和程序语言中，当前的线程都有一个独特的身份（id），并且这个身份信息可以以一个普通值的形式被很容易地获取到，典型的可以是一个integer或者指针值。

goroutine没有可以被程序员获取到的身份（id）的概念。



