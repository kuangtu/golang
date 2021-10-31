package main

import "fmt"



func main() {

    defer func() {
        fmt.Println("后执行")
        fmt.Println("正常退出")
    }()
    
    fmt.Println("开始执行")
    
    defer func() {
        v := recover()
        fmt.Println("先执行")
        fmt.Println("恐慌被恢复了:", v)
    }()
    
    panic("恐慌") //产生一个恐慌
    fmt.Println("无法执行")

}