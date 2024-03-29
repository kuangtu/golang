# 第零章  开场白

 笔者学习Go语言主要基于如下几个方面：

- 发明人中有C语言作者参与，笔者工作中主要开发语言是C语言，对于C语言的改进，或者结合目前互联网服务等环境，看起来Go语言更为适合；
- C++太复杂了；
- Go语言使用排名等在上升，国内一些大厂也在使用，生态环境逐渐丰富；
- 在证券行业中如交易所、广发证券等也都有使用；
- 在快速开发和执行速度之间架起了桥梁。

对于学习过程进行记录和分享。

## 设计哲学

### 简单

Go 语言的设计者们在语言设计之初，就拒绝了走语言特性融合的道路，选择了“做减法”并致力于打造一门简单的编程语言。

### 显式

在 Go 语言中，不同类型变量是不能在一起进行混合计算的，这是因为 Go 希望开发人员明确知道自己在做什么，这与 C 语言的“信任程序员”原则完全不同，因此你需要以显式的方式通过转型统一参与计算各个变量的类型。

### 组合

提供了正交的语法元素，供后续组合使用：

- Go 语言无类型层次体系，各类型之间是相互独立的，没有子类型的概念；
- 包之间相对独立，没有自包的概念；
- 实现某个接口时，无需像 Java 那样采用特定关键字修饰。

提供了“类型嵌入”，可以将已经实现的功能嵌入到新类型中，类似面向对象语言中的“继承”机制。且被嵌入的类型和新类型之间没有关系，甚至相互完全不知道对方的存在。

比如：

```go

// $GOROOT/src/sync/pool.go
type poolLocal struct {
    private interface{}   
    shared  []interface{}
    Mutex               
    pad     [128]byte  
}
```

poolLocal结构体中嵌入了类型Mutex，可以通过poolLocal类型的变量直接调用Mutex类型的Lock、Unlock方法。

### 并发

Go 放弃了传统的基于操作系统线程的并发模型，而采用了用户层轻量级线程，Go 将之称为 goroutine。

goroutine 占用的资源非常小，Go 运行时默认为每个 goroutine 分配的栈空间仅 2KB。goroutine 调度的切换也不用陷入（trap）操作系统内核层完成，代价很低。

且内置了了辅助并发设计的原语：channel和select。开发者可以通过语言内置的 channel 传递消息或实现同步，并通过 select 实现多路 channel 的并发控制。

### 面向工程

Go 语言设计的初衷，就是面向解决真实世界中 Google 内部大规模软件开发存在的各种问题，为这些问题提供答案，这些问题包括：程序构建慢、依赖管理失控、代码难于理解、跨语言构建难等。

- 重新设计编译单元和目标文件格式，实现 Go 源码快速构建，让大工程的构建时间缩短到类似动态语言的交互式解释的编译速度；
- 如果源文件导入它不使用的包，则程序将无法编译，不会编译额外的代码，缩短编译时间；
- 包路径是唯一的，而包名不必唯一的，降低包名起名的负担；
- 标准库丰富，其中net/http、crypto、encoding等包充分迎合了云原生时代的关于 API/RPC Web 服务的构建需求。



## Go语言特性

- Go语言定义了能做什么，还定义了不能做什么；
- 语法简洁，便于记忆；
- Go语言编译器速度非常快，显著减少项目构建的时间；
- 内置并发机制，不用使用特定的线程库，即：*go*routine机制；
- 自带垃圾回收机制，不需要用户自己管理内存。

## 开发速度

Go 语言使用了更加智能的编译器，并简化了解决依赖的算法，最终提供了更快的编译速度。 编译 Go 程序时，编译器只会关注那些直接被引用的库，而不是像 Java、C 和 C++那样，要遍历依赖链中所有依赖的库。

**C/C++编译**： 编译时-l -L选项遍历依赖的库。







