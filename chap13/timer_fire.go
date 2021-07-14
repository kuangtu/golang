package main

import (
	"fmt"
	"time"
)

func ExecFunc() {

	fmt.Println("execute function")
}

func main() {
	for {
		fmt.Println("create timer")
		timer1 := time.NewTimer(2 * time.Second)

		<-timer1.C
		fmt.Println("timer 1fired")
		ExecFunc()
	}

}
