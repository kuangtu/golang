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
    addr, err := net.ResolveIPAddr(name)
    if err != nil {
        fmt.Println("Resolved IPAddr error", err.Error())
        os.Exit(1)
    }
    fmt.Println("Resolved IPAddr is", addr.String()) 
    
    os.Exit(0)
}
    