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
	//消息字符串填充
	SenderCompID     = "CSI"
	TargetCompID     = "SSE"
	LOGIN_MSGTYPE    = "S001"
	MsgType_LEN      = 4
	SenderCompID_LEN = 32
	TargetCompID_LEN = 32
	AppVerID_LEN     = 8
	//消息体头部长度
	MsgHeader_LEN = 24
	//消息体长度
	LoginMsg_BODY_LEN = 15
	LOGINMSG_TAIL_LEN = 4
)

var (
	Pkt_MsgSeq uint64 = 1
	HeartBtInt uint16 = 10
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

func initLoginMsg(loginMsg *LoginMsg) {

	//按照接口规范初始化char字符串类型，通过空格填充
	for i, c := range loginMsg.SenderCompID {
		loginMsg.SenderCompID[i] = ' '
	}

	for i, c := range loginMsg.TargetCompID {
		loginMsg.TargetCompID[i] = ' '
	}

	for i, c := range loginMsg.AppVerID {
		loginMsg.AppVerID = ' '
	}

}

func setLoginMsgBody(loginMsg *LoginMsg) {
	//填充发送ID
	var setStr []byte
	setStr = []byte(SenderCompID)
	for i, c := range setStr {
		loginMsg.SenderCompID[i] = c
	}

	//填充目标ID
	setStr = []byte(TargetCompID)
	for i, c := range setStr {
		loginMsg.TargetCompID[i] = c
	}

	//填充APP
	setStr = []byte(AppVerID)
	for i, c := range setStr {
		loginMsg.AppVerID[i] = c
	}

	//填充心跳周期
	loginMsg.HeartBtInt = HeartBtInt
}

func setLoginMsgHeader(loginMsg *LoginMsg) {

}

func main() {
	//创建登录消息
	loginMsg := LoginMsg{}

	//消息初始化
	initLoginMsg(&loginMsg)
	fmt.Printf("%v", loginMsg)

	//填充字段
	setLoginMsgBody(&loginMsg)

	//填充消息头部

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
