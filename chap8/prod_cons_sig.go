package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

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
