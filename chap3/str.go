package main

import (
    "fmt"
)

func main() {
    s1 := "hello world"
    fmt.Println("the str len is:", len(s1))
    fmt.Printf("the second char is:0x%x", s1[1])
}