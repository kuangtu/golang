package main

import(
    "net"
    "os"
    "fmt"
    "time"
)

func main() {

    service := ":1200"

    //解析TCP地址及端口
    tcpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)
    
    //监听端口
    listener, err := net.ListenUDP("udp", tcpAddr)
    checkError(err)
    
    for {
        handleClient(conn)
    }
    
}

func handleClient(conn *net.UDPConn) {
    var buf [512]byte
    _, addr, err := conn.ReadFromUDP(buf[0:])
    
    if err != nil {
        return
    }
    
    daytime := time.Now().String()
    
    conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
