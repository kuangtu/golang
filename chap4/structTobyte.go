package main

import (
	"encoding/binary"
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

var (
	Pkt_MsgSeq uint64 = 1
	HeartBtInt uint16 = 10
)

type MsgHeader struct {
	MsgType      [MsgType_LEN]byte
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
	MsgTail
}

func initLoginMsg(loginMsg *LoginMsg) {

	//按照接口规范初始化char字符串类型，通过空格填充
	for i, _ := range loginMsg.SenderCompID {
		loginMsg.SenderCompID[i] = ' '
	}

	for i, _ := range loginMsg.TargetCompID {
		loginMsg.TargetCompID[i] = ' '
	}

	for i, _ := range loginMsg.AppVerID {
		loginMsg.AppVerID[i] = ' '
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
	//填充消息类型
	var setStr []byte
	setStr = []byte(LOGIN_MSGTYPE)
	for i, c := range setStr {
		loginMsg.MsgType[i] = c
	}

	//填充消息序号
	loginMsg.MsgSeq = Pkt_MsgSeq
	Pkt_MsgSeq = Pkt_MsgSeq + 1

	//填充发送时间
	//获取当前时间
	loginMsg.SendingTtime = 1024

	//消息体长度
	loginMsg.BodyLength = LOGINMSG_BODY_LEN
}

func setLoginMsgTail(loginMsg *LoginMsg, chksum uint32) {
	//填充消息消息尾部校验码
	loginMsg.CheckSum = chksum
}

func doCalChkSum(buffer []byte, len uint32) uint32 {
	var chkSum uint8
	var i uint32
	for i = 0; i < len; i++ {
		chkSum += uint8(buffer[i])
	}

	return uint32(chkSum)
}

func calLogMsgChkSum(loginMsg *LoginMsg) ([]byte, uint32) {
	//创建byte slicetmp
	var slicetmp []byte
	b := make([]byte, MSGHEADER_LEN+LOGINMSG_BODY_LEN+LOGINMSG_TAIL_LEN)
	//消息头部放入到buffer中
	//消息类型
	slicetmp = loginMsg.MsgType[:]
	copy(b, slicetmp)
	//消息时间
	binary.BigEndian.PutUint64(b[MsgType_LEN:], loginMsg.SendingTtime)
	//消息序号
	binary.BigEndian.PutUint64(b[(MsgType_LEN+UINT64_LEN):], loginMsg.MsgSeq)
	//消息体长度
	binary.BigEndian.PutUint32(b[(MsgType_LEN+UINT64_LEN+UINT64_LEN):], loginMsg.BodyLength)

	//消息体写入
	//发送ID
	slicetmp = loginMsg.SenderCompID[:]
	copy(b[(MsgType_LEN+UINT64_LEN+UINT64_LEN+UINT32_LEN):], slicetmp)
	//目标ID
	slicetmp = loginMsg.TargetCompID[:]
	copy(b[(MsgType_LEN+UINT64_LEN+UINT64_LEN+UINT32_LEN+SenderCompID_LEN):], slicetmp)
	//心跳间隔
	binary.BigEndian.PutUint16(
		b[(MsgType_LEN+UINT64_LEN+UINT64_LEN+UINT32_LEN+SenderCompID_LEN+TargetCompID_LEN):],
		loginMsg.HeartBtInt)
	//APPver
	slicetmp = loginMsg.AppVerID[:]
	copy(b[(MsgType_LEN+UINT64_LEN+UINT64_LEN+UINT32_LEN+SenderCompID_LEN+TargetCompID_LEN+UINT16_LEN):], slicetmp)

	//计算校验码
	chksum := doCalChkSum(b, MSGHEADER_LEN+LOGINMSG_BODY_LEN)
	//校验码写入到buffer中
	binary.BigEndian.PutUint32(b[(MSGHEADER_LEN+LOGINMSG_BODY_LEN):], chksum)

	return b, chksum

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
	setLoginMsgHeader(&loginMsg)

	//计算校验码，并放入到byte数组中用于socket发送
	buffer, chksum := calLogMsgChkSum(&loginMsg)

	fmt.Println("the checksum is:", chksum)
	fmt.Print("the packet buffer is:", buffer)
}
