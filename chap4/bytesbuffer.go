package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	//消息字符串填充
	SenderCompID     = "CSI"
	TargetCompID     = "SSE"
	LOGIN_MSGTYPE    = "S001"
	AppVerID         = "1.00"
	MsgType_LEN      = 4
	SenderCompID_LEN = 32
	TargetCompID_LEN = 32
	AppVerID_LEN     = 8
	//消息体头部长度
	MSGHEADER_LEN = 24
	//消息体长度
	LOGINMSG_BODY_LEN = 74
	LOGINMSG_TAIL_LEN = 4
	UINT64_LEN        = 8
	UINT32_LEN        = 4
	UINT16_LEN        = 2
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
	SenderCompID [3]byte
	TargetCompID [3]byte
	HeartBtInt   uint16
	AppVerID     [4]byte
}

func main() {
	//初始化消息
	var loginMsg LoginMsg
	loginMsg.MsgType = [4]byte{'S', '0', '0', '1'}
	loginMsg.SendingTtime = 1
	loginMsg.MsgSeq = 2
	loginMsg.BodyLength = 12
	loginMsg.SenderCompID = [3]byte{'M', 'M', 'M'}
	loginMsg.TargetCompID = [3]byte{'S', 'S', 'S'}
	loginMsg.HeartBtInt = 3
	loginMsg.AppVerID = [4]byte{'1', '.', '0', '0'}

	buf := new(bytes.Buffer)
	//写入消息类型，
	buf.Write(loginMsg.MsgType[:])
	//按照大端写入发送时间
	binary.Write(buf, binary.BigEndian, loginMsg.SendingTtime)
	//写入消息序号
	binary.Write(buf, binary.BigEndian, loginMsg.MsgSeq)
	//写入消息体长度
	binary.Write(buf, binary.BigEndian, loginMsg.BodyLength)
	//写入发送ID
	buf.Write(loginMsg.SenderCompID[:])
	//写入目标ID
	buf.Write(loginMsg.TargetCompID[:])
	//写入心跳时间
	binary.Write(buf, binary.BigEndian, loginMsg.HeartBtInt)
	//写入版本信息
	buf.Write(loginMsg.AppVerID[:])

	//将这些byte数据写入到文件中
	f, err := os.Create("struct.bin")
	if err != nil {
		fmt.Println("open file error.")
	}
	defer f.Close()
	n, err := f.Write(buf.Bytes())
	if err != nil {
		fmt.Println("write file failed.")
	}
	fmt.Println("write file len is:", n)

}
