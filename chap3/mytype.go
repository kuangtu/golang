package main

import (
    "fmt"
)

type MyInt int

func main() {
    var a ,b c MyInt = 1, 2, 3
    c := a + b
    fmt.Println("the sum is:%d.", c)
}
