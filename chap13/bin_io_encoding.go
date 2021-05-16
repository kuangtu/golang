package main

import (
    "encoding/binary"
    "fmt"
    _ "time"
    "bytes"
)

type packet struct {
    Sensid uint32
    Locid  uint16
    Tstamp uint32
    Temp   int16
    cs string
}

func main() {
    dataIn := packet {
        Sensid: 1, Locid: 1233, Tstamp: 123452123, Temp: 12,cs: "test",
    }
    
    buf := new(bytes.Buffer)
    
    err := binary.Write(buf, binary.BigEndian, dataIn)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    //Buffer长度
    fmt.Printf("%d\n", buf.Len())
    
}