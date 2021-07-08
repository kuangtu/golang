# 第17章 内存模型



## 17.1 介绍

Go 内存模型指定了在何种条件下可以保证在一个 goroutine 中读取变量时观察到在不同 goroutine 中写入相同变量所产生的值。

多个goroutine 同时修改数据时必须串行化访问。为了能够序列化访问，通过管道或者其他同步原语比如：sync等保护数据。



## 17.2 Happens Before

在单个goroutine 内部，读和写必须按照程序指定的顺序执行。也就是说，当重排不会修改语言规范定义的goroutine中的行为时，编译器和处理器可能对单个goroutine 内的读和写进行重排。由于这种重新排序，一个 goroutine 观察到的执行顺序可能与另一个 goroutine 感知的顺序不同。比如：一个goroutine执行```a = 1; b = 2;```，个另一个可能会在更新a之前观察到 b的更新。

为了指定读写要求，在Go语言中定义了“*happens before*"，内存操作执行的偏序。

如果事件e1在e2之前发生，我们就说e2在e1之后发生；如果e1没有在e2之前发生也没有在e2之后发生，我们就说e1和e2同时发生的。

在单个goroutine内，happens-before顺序就是程序表达的顺序。

如果r读取变量v时能够观察到w写v，需要满足以下两个条件：

- r没有在w之前发生；
- 没有其他w'写v发生在w之后，r之前。

为了保证对变量v的读取r观察到对v的特定写入w，需要确保允许r唯一观察的写是w，需要满足以下两个条件：

- w 在r之前发生；
- 其他任何向变量v的写，发生在w之前，或者r之后。

在单个goroutine内部，没有并发，所以两个定义是等价的：一个读r观察最近一次写w到v写入的值。当多个goroutine访问一个共享变量v时，它们必须使用同步事件来建立happens-before 确保读取观察所需写入的条件。

## 17.3 同步（Synchronization）

 ### 17.3.1 初始化

程序初始化在单个的goroutine内运行，但该goroutine可能会创建其他并发运行的goroutine。

如果包 p 导入包 q，则 q 的 init 函数的完成发生在任何 p 的开始之前。

main 函数的启动发生在所有 init 函数完成之后。

### 17.3.2 创建goroutine

启动新 goroutine 的 go 语句发生在 goroutine 开始执行之前。

```
var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}
```





https://golang.org/ref/mem