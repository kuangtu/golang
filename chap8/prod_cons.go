package main

import (
	"fmt"
	"time"
)

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
