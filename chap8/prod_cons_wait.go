package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(queue chan string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("producer string item:", (i + 1))
		queue <- fmt.Sprintf("item %d", (i + 1))
		time.Sleep(time.Second * 10)
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
