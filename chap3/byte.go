package main

import (
    "fmt"
)

func BytesToString(data []byte) string{
    return string(data[:])
}

func main() {
    
    a := []byte{0,1,2,3,4}
    fmt.Println(a)
    stra := BytesToString(a)
    //因为是不可显示的，因此输出为
    fmt.Println(stra)
    
    b := []byte{48,49,50,51,52}
    fmt.Println(b)
    strb := BytesToString(b)
    fmt.Println(strb)
}