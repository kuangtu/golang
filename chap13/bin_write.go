package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type MsgHeader struct {
	MsgType      [4]byte
	SendingTtime uint64
	MsgSeq       uint64
	BodyLength   uint32
}

type MsgTail struct {
	CheckSum uint32
}

//登录消息
type LoginMsg struct {
	MsgHeader
	SenderCompID [32]byte
	TargetCompID [32]byte
	HeartBtInt   uint16
	AppVerID     [8]byte
	MsgTail
}

//初始化结构体消息
func initMsg(loginMsg *LoginMsg) {
	loginMsg.CheckSum = 512
	loginMsg.SendingTtime = 1
}

func main() {
	loginMsg := &LoginMsg{}
	initMsg(loginMsg)

	buf := new(bytes.Buffer)
	// binary.Write(buf, binary.BigEndian, loginMsg)
	binary.Write(buf, binary.LittleEndian, loginMsg)
	fmt.Println("the binary buf size:", buf.Len())
	fmt.Printf("%x", buf)

}
