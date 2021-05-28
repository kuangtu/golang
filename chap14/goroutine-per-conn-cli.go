package main

import (
    "net"
    "fmt"
    "flag"
    "time"
)

var (
    ip = flag.String("ip", "127.0.0.1", "Server IP")
    connections = flag.Int("conn", 1, "number of tcp connections")
)

func main() {
    //分析参数
    flag.Parse()
    
    addr := *ip + ":12345"
    fmt.Println("连接到:%s", addr)
    var conns []net.Conn
    
    for i := 0; i < *connections; i++   {
        c, err := net.DialTimeout("tcp", addr, 10 *time.Second)
        if err != nil {
            fmt.Println("failed to connect", i, err)
        }
        
        conns = append(conns, c)
        
        time.Sleep(time.Millisecond)
    }
    
    defer func() {
        for _, c := range conns {
            c.Close()
        }
    }()
    
    fmt.Println("完成初始化 %d 连接", len(conns))
    
    tts := time.Second
    
    if *connections > 100 {
        tts = time.Millisecond * 5
    }
    
    for  {
        for i := 0; i < len(conns); i++ {
            time.Sleep(tts)
            conn := conns[i]
            conn.Write([]byte("hello world\r\n"))
        }
    }
    
}




        
        