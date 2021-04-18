package main

import(
    "net"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "usage: %s ip-addr\n", os.Args[0])
        os.Exit(1)
    }
    
    name := os.Args[1]
    
    //解析IP地址
    addr := net.ParseIP(name)
    if addr == nil {
        fmt.Println("invalid address")
    }
    else {
        fmt.Println("the addr is:", addr.String())
    }
    
    os.Exit(0)
}
    