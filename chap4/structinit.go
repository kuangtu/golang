package main

import (
	"bytes"
	"fmt"
)

const (
	MsgType_LEN      = 4
	SenderCompID_LEN = 32
	TargetCompID_LEN = 32
	AppVerID_LEN     = 8
	//消息体长度
	MsgHeader_LEN     = 24
	LoginMsg_BODY_LEN = 15
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

type LoginMsg struct {
	MsgHeader
	SenderCompID [SenderCompID_LEN]byte
	TargetCompID [TargetCompID_LEN]byte
	HeartBtInt   uint16
	AppVerID     [AppVerID_LEN]byte
	CheckSum     uint32
	MsgTail
}

func putBuffer(i interface{}) (b bytes.Buffer, chksum uint32) {
	var slice1 []byte
	switch v := i.(type) {
	case LoginMsg:
		//fmt.Printf("it's login msg:%v", v)
		fmt.Printf("before it's a loginmsg")
		slice1 = v.MsgType[:]
		fmt.Printf("after it's a loginmsg")
		b.Write(slice1)
		// fmt.Printf(string(b.Bytes()))
		chksum = 5
	}

	return
}

func main() {

	//结构体中，整数、数组默认0值初始化
	var loginMsg = LoginMsg{}
	fmt.Println("the bodyLength is:", loginMsg.BodyLength)
	fmt.Println("the bytes is:", loginMsg.MsgType)
	//填充send id
	for i := range loginMsg.SenderCompID {
		loginMsg.SenderCompID[i] = ' '
	}
	var slice1 []byte
	slice1 = loginMsg.SenderCompID[:]
	copy(slice1, "a")
	fmt.Println(loginMsg.SenderCompID)

	b, chksum := putBuffer(loginMsg)
	fmt.Println(chksum)
	fmt.Println(b)

}
