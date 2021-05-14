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


type Foo21 struct {
    aaa int32 // 4
    bbb int32 // 4
    ccc int32 // 4
}

type Foo22 struct {
    aaa bool // 1
    bbb int32 // 4 (max)
    ссс bool // 1 
    ddd bool // 1
}

type Foo23 struct {
    aaa bool // 1
    ссс bool // 1 
    ddd bool // 1
}


type Foo24 struct {
    aaa bool
    bbb int64
    ссс bool
    ddd bool
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
    
    y21 := Foo21{}
    fmt.Println("struct padding")
    fmt.Println(unsafe.Sizeof(y21))
    
    y22 := Foo22{}
    fmt.Println("struct padding")
    fmt.Println(unsafe.Sizeof(y22))
    
    y23 := Foo23{}
    fmt.Println("struct padding")
    fmt.Println(unsafe.Sizeof(y23))
    
    var y24 Foo24
    y24.aaa = false
    y24.bbb = 1
    //y24.ccc = false
    y24.ddd = false
    fmt.Println(unsafe.Offsetof(y24.bbb))
    fmt.Println(unsafe.Offsetof(y24.ddd))
    
}