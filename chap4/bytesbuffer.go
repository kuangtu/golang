package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func putBuffer(i interface{}) (b bytes.Buffer, chksum uint32) {
	b = bytes.Buffer{}
	chksum = 1
	return b, chksum
}

func main() {
	var buf = bytes.Buffer{}
	buf.Write([]byte{'1', '2', '3'})
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
	fmt.Println(buf.Cap())
	mybytes := [3]byte{1, 2, 3}
	//基于切片的方式
	slice1 := mybytes[:]
	buf.Write(slice1)
	fmt.Println(buf.Bytes())
	fmt.Println(buf.Len())
	fmt.Println(buf.Cap())

	//写入整数
	// var buf1 = bytes.Buffer{}
	var seq uint32 = 3
	b := make([]byte, 9)
	//写入MsgType， MD001
	copy(b, "MD001")
	binary.BigEndian.PutUint32(b[5:], seq)
	//大端方式写入，最先读到的是高字节位
	fmt.Println(b)
	var sum int32 = 0
	for i := 0; i < 9; i++ {
		sum += (int32)(b[i])
	}
	fmt.Println("the sum is:", sum)
}
