package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("ticker at", t)
			}
		}
	}()
	fmt.Println("before sleep")
	//main goroutine等待5秒钟，然后ticker.Stop
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
