package main

import (
    "fmt"
    "unsafe"
)



func main() {
    var s string
    var c complex128
    fmt.Println(unsafe.Sizeof(s))   
    fmt.Println(unsafe.Sizeof(c))   
    
    var a [3]uint32
    fmt.Println(unsafe.Sizeof(a)) // prints 12
    
    
    var d struct{}
    fmt.Println(unsafe.Sizeof(d)) // prints 0
    
    type S struct {
        A struct{}
        B struct{}
    }
    var s2 S
    fmt.Println(unsafe.Sizeof(s2)) // prints 0
    
    var x [1000000000]struct{}
    fmt.Println(unsafe.Sizeof(x)) // prints 0


    var x1 = make([]struct{}, 1000000000)
    fmt.Println(unsafe.Sizeof(x1)) // prints 24 in the playground
}