package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func intDataSize(data interface{}) int {
	switch data := data.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []bool:
		return len(data)
	case []int8:
		return len(data)
	case []uint8:
		return len(data)
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return 2 * len(data)
	case []uint16:
		return 2 * len(data)
	case int32, uint32, *int32, *uint32:
		return 4
	case []int32:
		return 4 * len(data)
	case []uint32:
		return 4 * len(data)
	case int64, uint64, *int64, *uint64:
		return 8
	case []int64:
		return 8 * len(data)
	case []uint64:
		return 8 * len(data)
	case float32, *float32:
		return 4
	case float64, *float64:
		return 8
	case []float32:
		return 4 * len(data)
	case []float64:
		return 8 * len(data)
	}
	return 0
}

type data struct {
	A      uint32
	B      uint16
	Mytype [3]byte
}

func main() {
	b := []byte{0x01, 0x02, 0x02, 0x02, 0x02, 0x03, 0x01, 0x02, 0x03}
	buf := bytes.NewReader(b)
	fmt.Println("the reader len is:", buf.Len())

	var mydata data
	//直接读取到结构体中
	err := binary.Read(buf, binary.BigEndian, &mydata)
	if err != nil {
		fmt.Println("read data structure err")
	}

	fmt.Printf("%x\n", mydata.A)
	fmt.Printf("%x\n", mydata.B)
	fmt.Printf("%x\n", mydata.Mytype)
	fmt.Println("the reader len is:", buf.Len())

}
