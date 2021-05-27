package main

import (
    "net"
    "fmt"
    "io"
    "io/ioutil"
    "flag"
    "log"
    "time"
)

var (
    ip := flag.String("ip", "127.0.0.1", "Server IP")
    connections := flag.Int("conn", 1, "number of tcp connections")
)

func main() {
    //分析参数
    flag.Parse()
    
    addr := *ip + ":12345"
    fmt.Println("连接到:%s", addr)
    var conns []net.Conn
    
    for i := 0; i < *connections; i++   {
        c, err := net.DailTimeout("tcp", addr, 10 *time.Second)
        if err != nil {
            fmt.Println("failed to connect", i, err)
        }
        
        conns = append(conns, c)
        
        time.Sleep(time.Millisecond)
    }
    
    def func() {
        for _, c := range conns {
            c.Close()
        }
    }()
    
    fmt
        
    sock, err := net.Listen("tcp", "12345")
    if err != nil {
        return
    }
    
    defer func() {
        for _, conn := range connections {
            conn.Close()
        }
    }()
    
    for {
        conn, e := sock.Accept()
        if e != nil {
            if ne, ok := e.(net.Error); ok && ne.Temporary() {

                continue
            }
        }
        
        go handleConn(conn)
        
        connections = append(connections, conn)
    }
    
}

func handleConn(conn net.Conn) {
    io.Copy(ioutil.Discard, conn)
    fmt.Println("handleConn")
}



        
        