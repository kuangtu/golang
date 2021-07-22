package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	msgQue = make(chan int, 10)
	quit   = make(chan bool)
	wait   sync.WaitGroup
	ticker = time.NewTicker(time.Second * 1)
)

func Proc() {
	var iCnt int
	for {
		fmt.Println("start proc goroutine")
		msgQue <- 1
		time.Sleep(time.Second * 5)
		iCnt += 1
		if iCnt == 5 {
			msgQue <- 0
			break
		}
	}
	wait.Done()
	//通知心跳消息goroutine结束
	fmt.Println("send heart quit channel")
	quit <- true
}

func Cons() {
	for {
		fmt.Println("start consumer goroutine")
		a := <-msgQue
		fmt.Println("get: ", a)
		if a == 0 {
			break
		}
	}
	wait.Done()

}

func HeartBt() {
	for {
		fmt.Println("heart beat for")

		select {
		case <-quit:
			fmt.Println("proc and cons finished, heartbeat quit.")
			return
		case t := <-ticker.C:
			fmt.Println("heart beat current time: ", t)
		}
	}

}

func main() {
	defer ticker.Stop()
	//创建三个goroutine，
	//（1）模拟发送心跳消息
	//（2）模拟接收行情
	//（3）模拟处理行情

	wait.Add(2)
	go Proc()
	go Cons()
	go HeartBt()
	wait.Wait()
	fmt.Println("main finished.")
}
