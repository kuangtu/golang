package main

import (
    "net"
    "fmt"
    "io"
    "io/ioutil"
)

func main() {
    var connections []net.Conn
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



        
        