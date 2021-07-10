# 第8章 Goroutines和Channels

并发程序指同时进行多个任务的程序，Go语言中的并发程序可以用两种手段来实现。本章讲解goroutine和channel，其支持“顺序通信进程”（communicating sequential processes）或被简称为CSP。

Go 语言的并发同步模型来自一个叫作通信顺序进程（Communicating Sequential Processes，CSP） 的范型（paradigm）。CSP 是一种消息传递模型，通过在 goroutine 之间传递数据来传递消息，而不是对数据进行加锁来实现同步访问。用于在 goroutine 之间同步和传递数据的关键数据类型叫作通道 （channel）。



## 8.1 并发与并行

进程可以看做包含了应用程序在运行中需要用到和维护的各种资源的容器。这些资源包括但不限于内存地址空 间、文件和设备的句柄以及线程。

一个线程是一个执行空间，这 个 空间会被操作系统调度来运行 函数中所写的代码。每个进程至少包含一个线程，每个进程的初始线程被称作主线程。主线程终止应用程序也终止。

![进程线程描绘](C:\moxuansheng\workspace\golang\jpg\进程线程描绘.png)

操作系统会在物理处理器上调度线程来运行，而 Go 语言的运行时会在逻辑处理器上调度 goroutine来运行。每个逻辑处理器都分别绑定到单个操作系统线程。**协程更小**

![go调度器管理协程](C:\moxuansheng\workspace\golang\jpg\go调度器管理协程.png)

并发（concurrency）不是并行（parallelism）。并行是让不同的代码片段同时在不同的物理处 理器上执行。并行的关键是同时做很多事情，而并发是指同时管理很多事情，这些事情可能只做 了一半就被暂停去做别的事情了。在很多情况下，并发的效果比并行好，因为操作系统和硬件的 总资源一般很少，但能支持系统同时做很多事情。这种“使用较少的资源做更多的事情”的哲学， 也是指导 Go 语言设计的哲学。



## 8.2 Goroutines

每一个并发的执行单元叫作一个goroutine。简单地把goroutine类比作一个线程。**但是有本质的区别**



当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

```go
f()    // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

然后主函数返回。主函数返回时，所有的goroutine都会被直接打断，程序退出。除了从主函数退出或者直接终止程序之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行。






## 8.2并发的Clock服务

网络编程是并发大显身手的一个领域，由于服务器是最典型的需要同时处理很多连接的程序，这些连接一般来自于彼此独立的客户端。

如下是一个顺序执行的时钟服务器：

```go
// Clock1 is a TCP server that periodically writes the time.
package main

import (
    "io"
    "log"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err) // e.g., connection aborted
            continue
        }
        handleConn(conn) // handle one connection at a time
    }
}

func handleConn(c net.Conn) {
    defer c.Close()
    for {
        _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
        if err != nil {
            return // e.g., client disconnected
        }
        time.Sleep(1 * time.Second)
    }
}
```

每次服务一个客户端。

- Listen函数创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接，在这个例子里我们用的是TCP的localhost:8000端口。listener对象的Accept方法会直接阻塞，直到一个新的连接被创建，然后会返回一个net.Conn对象来表示这个连接。
- handleConn函数会处理一个完整的客户端连接。在一个for死循环中，用time.Now()获取当前时刻，然后写到客户端。由于net.Conn实现了io.Writer接口，我们可以直接向其写入内容。这个死循环会一直执行，直到写入失败。最可能的原因是客户端主动断开连接。这种情况下handleConn函数会用defer调用关闭服务器侧的连接，然后返回到主函数，继续等待下一个连接请求。
- **time.Time.Format方法提供了一种格式化日期和时间信息的方式。**



简单的netcat程序如下：

```go
// Netcat1 is a read-only TCP client.
package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
```

这个程序会从连接中读取数据，并将读到的内容写到标准输出中，直到遇到end of file的条件或者发生错误。

以上的程序示例只能同时服务一个客户端，第一个客户端完成之后才能接受第二个客户端的连接。

我们这里对服务端程序做一点小改动，使其支持并发：在handleConn函数调用的地方增加go关键字，让每一次handleConn的调用都进入一个独立的goroutine。

```go
for {
    conn, err := listener.Accept()
    if err != nil {
        log.Print(err) // e.g., connection aborted
        continue
    }
    go handleConn(conn) // handle connections concurrently
}
```

## 8.3 并发的Echo服务





## 8.4 Channels

在 Go 语言里，你不仅可以使用原子函数和互斥锁来保证对共享资源的安全访 问以及消除竞争状态，还可以使用通道，通过发送和接收需要共享的资源，在 goroutine 之间做 同步。



如果说goroutine是Go语言程序的并发体的话，那么channels则是它们之间的通信机制。一个channel是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。每个channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。

使用内置的make函数，我们可以创建一个channel：

```go
ch := make(chan int) // ch has type 'chan int'
```

和map类似，channel也对应一个make创建的底层数据结构的引用。当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel对象。和其它的引用类型一样，channel的零值也是nil。

```go
// 无缓冲的整型通道 
unbuffered := make(chan int)
// 有缓冲的字符串通道
buffered := make(chan string, 10)
```



一个channel有发送和接受两个主要操作，都是通信行为。一个发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine。发送和接收两个操作都使用`<-`运算符。在发送语句中，`<-`运算符分割channel和要发送的值。在接收语句中，`<-`运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。

```go
ch <- x  // a send statement 通过通道发送一个值
x = <-ch // a receive expression in an assignment statement 从通道接收一个值
<-ch     // a receive statement; result is discarded
```

### 无缓冲通道

无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通 道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。

![无缓冲通道同步](C:\moxuansheng\workspace\golang\jpg\无缓冲通道同步.png)

```go
// 这个示例程序展示如何用无缓冲的通道来模拟
// 4个  goroutine间的接力比赛
package main

import ( 
   "fmt"
   "sync"
   "time"
) 

// wg用来等待程序结束
var wg sync.WaitGroup

// main是所有  Go程序的入口
func main() { 
   // 创建一个无缓冲的通道
   baton := make(chan int)

   // 为最后一位跑步者将计数加  1 
   wg.Add(1)

   // 第一位跑步者持有接力棒
   go Runner(baton)

   // 开始比赛，模拟拿到了接力棒
   baton <- 1 

   // 等待比赛结束
   wg.Wait()
} 

// Runner模拟接力比赛中的一位跑步者
func Runner(baton chan int) { 
   var newRunner int
    // 等待接力棒
    runner := <-baton

    // 开始绕着跑道跑步
    fmt.Printf("Runner %d Running With Baton\n", runner)

    // 创建下一位跑步者
    if runner != 4 { 
        newRunner = runner + 1 
        fmt.Printf("Runner %d To The Line\n", newRunner)
        go Runner(baton)
    } 

    // 围绕跑道跑
    time.Sleep(100 * time.Millisecond)

    // 比赛结束了吗？
    if runner == 4 { 
        fmt.Printf("Runner %d Finished, Race Over\n", runner)
        wg.Done()
        return
    } 

    // 将接力棒交给下一位跑步者
    fmt.Printf("Runner %d Exchange With Runner %d\n",
        runner,
        newRunner)

    baton <- newRunner
}
```



### 有缓冲的通道

有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类 型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。通道会阻塞发送和接收动作的 条件也会不同。

只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲 区容纳被发送的值时，发送动作才会阻塞。



## 8.5 生产者消费者模式

可以通过channel方便地实现经典的生产者、消费者模式。

### 8.5.1 实现1

生产者将数据写入到管道，消费者从管道读取数据。

```go
func ProducerData(i int) int {
	fmt.Println("sleep..")
	time.Sleep(5 * time.Second)
	return i + 1
}
func main() {
	//创建管道
	data := make(chan int)

	//每隔5秒钟生产一个数据
	go func() {
		var i = 0
		for {
			i = ProducerData(i)
			data <- i
		}
	}()

	//消费者从管道中取出
	for i := range data {
		fmt.Printf("i=%v\n", i)
	}
}
```



匿名函数调用ProducerData函数，向管道中间隔接入数据。main函数中消费者从管道中读取。

### 8.5.2 实现2

也可以将生产者和消费者均通过goroutine启动，main主goroutine等待结束。

```go

func producer(queue chan string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("producer string item:", (i + 1))
		queue <- fmt.Sprintf("item %d", (i + 1))
		time.Sleep(time.Second)
	}

	//关闭channel
	close(queue)
	waitGroup.Done()
}

func consumer(queue chan string, waitGroup *sync.WaitGroup) {
	for val := range queue {
		fmt.Println("consuming the :", val)
	}

	waitGroup.Done()
}

func main() {
	fmt.Println("producer and consumer demo")
	queue := make(chan string)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go producer(queue, &waitGroup)
	go consumer(queue, &waitGroup)

	//等待goroutine结束
	waitGroup.Wait()
	fmt.Println("finished...")
}

```

### 8.5.3 实现3

生产者消费者模式运行过程中，如果碰到Ctrl-C信号程序直接退出。可以增加对于Ctrl-C操作系统信号的处理。

```go
func handleSigInt(sigInt chan os.Signal, queue chan string, quit chan bool) {
	_ = <-sigInt
	fmt.Println("recv singal and terminating gracefully")
	quit <- true
}

func producer(queue chan string, waitGroup *sync.WaitGroup, quit chan bool) {
	//退出是关闭了channel
	defer close(queue)
	for i := 0; i < 10; i++ {
		//不同的channel
		select {
		//退出
		case <-quit:
			waitGroup.Done()
			return
		default:
			fmt.Println("producing the item: ", (i + 1))
			queue <- fmt.Sprintf("item %d", (i + 1))
			time.Sleep(5 * time.Second)
		}

	}
	//消息生成之后
	waitGroup.Done()
}

func consumer(queue chan string, waitGroup *sync.WaitGroup) {
	for val := range queue {
		fmt.Println("consume the: ", val)
	}
	waitGroup.Done()
}

func main() {
	fmt.Println("producer consumer demo")
	queue := make(chan string)
	sigInt, quit := make(chan os.Signal), make(chan bool)
	//设置singal函数
	signal.Notify(sigInt, syscall.SIGINT, syscall.SIGTERM)

	go handleSigInt(sigInt, queue, quit)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go producer(queue, &waitGroup, quit)

	go consumer(queue, &waitGroup)

	//等待线程结束
	waitGroup.Wait()
	fmt.Println("fnished..")

}

```



## 8.6 goroutine执行顺序

当多个goroutine执行时，类似“多线程”方式，启动执行的顺序不确定。

### 8.6.1 函数顺序执行

假设三个function，每个执行3秒。

```go

func FirstFunc() {

	fmt.Println("start exec first function")
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec first function")

}

func SecondFunc() {
	fmt.Println("start exec second function")
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec second function")
}

func ThrdFunc() {
	fmt.Println("start exec third function")
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec third function")
}

func main() {
	fmt.Println("test")
	start := time.Now()

	fmt.Println(start.String())

	FirstFunc()

	SecondFunc()

	ThrdFunc()

	//time.Since,从之前的time到现在为止的时间。
	//time.Now().Sub(t)
	fmt.Printf("total time to finished: %s\n", time.Since(start).String())

}
```

顺序执行结果：

![goroutine_seq_exec_time](jpg/goroutine_seq_exec_time.png)

其中time包中```Now()```函数获取当前时间，```Since(time)```获取从time到目前的时间间隔。

###  8.6.2 并发执行

三个函数按照goroutine方式执行：

```go

func FirstFunc() {
	defer wg.Done()
	fmt.Println("start exec first function")
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec first function")

}

func SecondFunc() {
	defer wg.Done()
	fmt.Println("start exec second function")
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec second function")
}

func ThrdFunc() {
	defer wg.Done()
	fmt.Println("start exec third function")
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec third function")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("test")
	start := time.Now()
	wg.Add(3)
	fmt.Println(start.String())

	go FirstFunc()

	go SecondFunc()

	go ThrdFunc()

	//time.Since,从之前的time到现在为止的时间。
	//time.Now().Sub(t)
	wg.Wait()
	fmt.Printf("total time to finished: %s\n", time.Since(start).String())

}

```



需要注意的是：如果main方法中**，没有阻塞。其中的go不会执行。？**

所以通过waitGroup等待三个gorouine并发执行，结果为：

![goroutine_con_exec_time](jpg/goroutine_con_exec_time.png)

其中执行的顺序是不同的。

### 8.6.3 顺序并发执行1

在上一节中多个goroutine开始执行时，不保证顺序，可以通过channel通知方式，确定开始执行的顺序。

```go
func FirstFunc(ch chan int) {
	defer wg.Done()
	seq := <-ch
	fmt.Println("start exec first function:", seq)
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec first function")

}

func SecondFunc(ch chan int) {
	defer wg.Done()
	seq := <-ch
	fmt.Println("start exec sec function:", seq)
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec second function")
}

func ThrdFunc(ch chan int) {
	defer wg.Done()
	seq := <-ch
	fmt.Println("start exec third function:", seq)
	time.Sleep(3 * time.Second)
	fmt.Println("finish exec third function")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("test")
	start := time.Now()
	wg.Add(3)
	fmt.Println(start.String())

	//通过通道的方式保证goroutine执行顺序
	firstChan := make(chan int)
	secChan := make(chan int)
	thrdChan := make(chan int)

	go FirstFunc(firstChan)
	//按照顺序触发
	firstChan <- 1
	go SecondFunc(secChan)
	secChan <- 2
	go ThrdFunc(thrdChan)
	thrdChan <- 3

	//time.Since,从之前的time到现在为止的时间。
	//time.Now().Sub(t)
	wg.Wait()
	fmt.Printf("total time to finished: %s\n", time.Since(start).String())

}
```

创建无缓冲区管道，FirstFunc、SecondFunc和ThrdFunc函数在开始执行时，  ```seq := <-ch``` 从管道中读取int值。main函数中按照顺序向管道中写入。执行结果如下：

![goroutine_con_seq_exec_time](jpg/goroutine_con_seq_exec_time.png)

开始执行的顺序有了保障，但是完成执行的时间没有保障。

### 8.6.4 顺序并发执行2

在实际业务过程中会碰到一个场景，就多个goroutine执行，其中一个goroutineA中的部分“逻辑”完成之后，另外的一个goroutineB才能够B中的”逻辑“。比如：A接收数据然后放置到B中解析。考虑如下的场景：

（1）goroutineA启动网络接收处理逻辑，开始接收数据；

（2）goroutineB初始化buffer等，供A放入数据；

（3）如果goroutineB还没有初始化完成，goroutineA已经收到了数据并执行了“存放到buffer”中的逻辑，出现异常；

（4）goroutineB初始化“逻辑”需要在goroutineA存放数据之前完成。



https://medium.com/technofunnel/understanding-golang-and-goroutines-72ac3c9a014d





https://medium.com/@fonseka.live/how-to-maintain-the-order-of-go-routines-117a5be86c4f