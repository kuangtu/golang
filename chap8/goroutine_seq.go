package main

import (
	"fmt"
	"time"
)

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
