package main

import (
	"fmt"
	"sync"
	"time"
)

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
