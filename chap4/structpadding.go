package main

import (
    "fmt"
    "unsafe"
)

type Foo1 struct {

}

type Foo2 struct {
    a bool
    b int32
}

type Foo3 struct {
    a bool
    b int32
    c bool
    d bool
}

func main() {
    x1 := &Foo1{}
    y1 := Foo1{}
    fmt.Println("struct padding")
    fmt.Println(unsafe.Sizeof(x1))
    fmt.Println(unsafe.Sizeof(y1))
    
    var b bool
    var i int32
    fmt.Println("type sizeof")
    fmt.Println(unsafe.Sizeof(b))
    fmt.Println(unsafe.Sizeof(i))
    
    x2 := &Foo2{}
    y2 := Foo2{}
    fmt.Println("struct padding")
    fmt.Println(unsafe.Sizeof(x2))
    fmt.Println(unsafe.Sizeof(y2))
    
    x3 := &Foo3{}
    y3 := Foo3{}
    fmt.Println("struct padding")
    fmt.Println(unsafe.Sizeof(x3))
    fmt.Println(unsafe.Sizeof(y3))
    

    
}