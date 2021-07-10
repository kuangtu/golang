package main

import (
	"fmt"
	"sync"
	"time"
)

func ReadFunc(ch chan string) {

	fmt.Println("init socket")

	fmt.Println("wait for Proc Function init.")
	<-ch
	fmt.Println("read from socket")

	fmt.Println("write to ProcFunc buffer")

	wg.Done()
}

func ProcFunc(ch chan string) {

	fmt.Println("init buffer")

	ch <- "Proc init finished"

	fmt.Println("wait for data")

	wg.Done()
}

var wg sync.WaitGroup

func main() {
	fmt.Println("test")
	start := time.Now()
	wg.Add(2)

	notice := make(chan string)

	go ReadFunc(notice)

	go ProcFunc(notice)

	//time.Since,从之前的time到现在为止的时间。
	//time.Now().Sub(t)
	wg.Wait()
	fmt.Printf("total time to finished: %s\n", time.Since(start).String())

}
