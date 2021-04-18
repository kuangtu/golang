package main

import(
    "net"
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "usage: %s network-type serivce\n", os.Args[0])
        os.Exit(1)
    }
    
    netType := os.Args[1]
    service := os.Args[2]
    
    //查询端口服务
    port, err := net.LookupPort(netType, service)
    if err != nil {
        fmt.Println("Error:", err.Error())
        os.Exit(2)
    }
    fmt.Println("Service port:", port) 
    
    os.Exit(0)
}
    