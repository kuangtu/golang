package main

import (
	"fmt"
)

type LoginMsg struct {
	MsgType      [4]byte
	Time         uint64
	SenderCompID [32]byte
	TargetCompID [32]byte
	HeartBtInt   uint16
	AppVerID     [8]byte
	CheckSum     uint32
}

func main() {
	fmt.Println("test")
	b := []byte("M001")
	loginMsg := LoginMsg{Time: 111, HeartBtInt: 111}
	fmt.Println(loginMsg.Time)
	//通过数组赋值
	loginMsg.MsgType = [4]byte{'M', '0', '0', '1'}
	fmt.Println(loginMsg.MsgType)

	//通过切片赋值
	for i, x := range b {
		loginMsg.MsgType[i] = x
	}

}
