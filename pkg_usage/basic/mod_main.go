package main

import (
    "fmt"
    "container/list"
)

func main() {
    fmt.Println("create list")
    //beego.Run()
    l := list.New()
    l.PushBack(4)
    l.PushBack("H")
    
    //遍历
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}