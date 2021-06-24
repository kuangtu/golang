package main

import (
	"bytes"
	"fmt"
)

type CheckSum interface {
	CalCheckSum() uint32
}

func GetPktChkSum(chksum CheckSum) uint32 {
	return chksum.CalCheckSum()
}

const (
	MsgType_LEN      = 4
	SenderCompID_LEN = 32
	TargetCompID_LEN = 32
	AppVerID_LEN     = 8
	//消息体长度
	MsgHeader_LEN     = 24
	LoginMsg_BODY_LEN = 15
)

var (
	Pkt_MsgSeq uint64 = 1
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
	Tail         MsgTail
}

func putBuffer(i interface{}) (b bytes.Buffer, chksum uint32) {
	switch v := i.(type) {
	case LoginMsg:
		fmt.Println("it's login message", v.AppVerID)
		slicetmp := v.AppVerID[:]
		b.Write(slicetmp)
		//写入HeartBtInt

		// for index := 0; index < MsgType_LEN; index++ {
		// 	b.WriteByte(byte(v.MsgType[index]))

		// }
		chksum = 5
	}

	return
}

func main() {
	fmt.Println("test")
	var mystr []byte
	mystr = []byte("M001")
	loginMsg := LoginMsg{}

	//填充发送ID
	var index int
	for index = 0; index < SenderCompID_LEN; index++ {
		loginMsg.SenderCompID[index] = ' '
	}
	mystr = []byte("csisender")
	for i, x := range mystr {
		loginMsg.SenderCompID[i] = x
	}

	//填充目标ID
	for index = 0; index < TargetCompID_LEN; index++ {
		loginMsg.TargetCompID[index] = ' '
	}
	mystr = []byte("csitarget")
	for i, x := range mystr {
		loginMsg.TargetCompID[i] = x
	}

	//填充APPVerID
	for index = 0; index < AppVerID_LEN; index++ {
		loginMsg.AppVerID[index] = ' '
	}
	mystr = []byte("1.00")
	for i, x := range mystr {
		loginMsg.AppVerID[i] = x
	}

	//填充消息头部
	loginMsg.MsgType = [4]byte{'M', '0', '0', '1'}
	loginMsg.SendingTtime = 1234
	loginMsg.MsgSeq = Pkt_MsgSeq
	loginMsg.BodyLength = LoginMsg_BODY_LEN

	fmt.Println("the body len is:", loginMsg.BodyLength)
	//计算消息的CHECKSUM

	//将结构体放入到byteBuferr中
	b, chksum := putBuffer(loginMsg)
	fmt.Println(b.String())
	fmt.Println(chksum)
}
