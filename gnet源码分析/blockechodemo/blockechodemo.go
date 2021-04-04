package main

import (
    "log"
    "time"

    "github.com/panjf2000/gnet"
    "github.com/panjf2000/gnet/pool/goroutine"
)

type echoServer struct {
    //嵌入类型
    *gnet.EventServer
    pool *goroutine.Pool
}

func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
    //增加数据
    data := append([]byte{}, frame...)
    //通过独立的协程运行
    // Use ants pool to unblock the event-loop.
    _ = es.pool.Submit(func() {
        //业务逻辑中有阻塞代码。放入到协程中运行
        time.Sleep(1 * time.Second)
        c.AsyncWrite(data)
    })

    return
}

func main() {
    p := goroutine.Default()
    defer p.Release()

    echo := &echoServer{pool: p}
    log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}