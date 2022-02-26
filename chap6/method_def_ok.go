package main

import (
    "fmt"
    _ "io"
)


type T struct{}

func (t T) M(n int) {
    fmt.Println(n)
}

func main() {
    var t T
    t.M(1) // 通过类型T的变量实例调用方法M

    p := &T{}
    p.M(2) // 通过类型*T的变量实例调用方法M
}