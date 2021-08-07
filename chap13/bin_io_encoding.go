package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	_ "time"
)

type packet struct {
	Sensid uint32
	Locid  uint16
	Tstamp uint32
	Temp   int16
}

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
	Header       MsgHeader
	SenderCompID [32]byte
	TargetCompID [32]byte
	HeartBtInt   uint16
	AppVerID     [8]byte
	Tail         MsgTail
}

func main() {
	dataIn := packet{
		Sensid: 1, Locid: 1233, Tstamp: 123452123, Temp: 12,
	}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, dataIn)

	if err != nil {
		fmt.Println(err)
		return
	}

	//Buffer长度
	fmt.Printf("%d\n", buf.Len())

	var dataOut packet
	if err := binary.Read(buf, binary.BigEndian, &dataOut); err != nil {
		fmt.Println("failed to Read:", err)
		return
	}

	fmt.Printf("the Sensid is:%d\n", dataOut.Sensid)

	//消息头部写入
	buf.Reset()

	header := MsgHeader{
		MsgType: [...]byte{0, 0, 0, 0}, SendingTtime: 1, MsgSeq: 1, BodyLength: 22,
	}

	tailmsg := MsgTail{
		CheckSum: 0x11,
	}

	loginmsg := LoginMsg{
		Header: header, Tail: tailmsg,
	}

	err = binary.Write(buf, binary.BigEndian, loginmsg)

	if err != nil {
		fmt.Println(err)
		return
	}

}
