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
    
    //解析查询IP地址
    addrs, err := net.LookupHost(name)
    if err != nil {
        fmt.Println("Resolved IPAddr error", err.Error())
        os.Exit(2)
    }
    
    for _, s := range addrs {
        fmt.Println(s)
    }
    
    os.Exit(0)
}
    