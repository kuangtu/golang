package main

import (
	"fmt"
	"sync"
	"time"
)

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
