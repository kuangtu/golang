package main

import (
    "encoding/binary"
    "fmt"
    _ "time"
)

func main() {
    buf := make([]byte, 10)
    //ts := uint32(time.Now().Unix())
    
    binary.BigEndian.PutUint16(buf[0:], 0x0101)
    binary.BigEndian.PutUint16(buf[2:], 0x0201)
    binary.BigEndian.PutUint32(buf[4:], 0x0301)
    binary.BigEndian.PutUint16(buf[8:], 0x0401)
    
    fmt.Printf("%x\n", buf)
    
    sensorID := binary.BigEndian.Uint16(buf[0:])
    locID :=  binary.BigEndian.Uint16(buf[2:])
    tstamp :=  binary.BigEndian.Uint32(buf[4:])
    temp :=  binary.BigEndian.Uint16(buf[8:])
    
    fmt.Printf("sid: %0#x, locID %0#x ts: %0#x, temp:%d\n", sensorID, locID, tstamp, temp)
}