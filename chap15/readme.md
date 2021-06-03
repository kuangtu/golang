# 第15章 web编程

本章对使用golang进行web编程进行整理。



## 15.1 概述

HTTP协议是互联网的基础协议，也是网页开发必备的基础知识，[可以参照阮一峰老师的文章](http://www.ruanyifeng.com/blog/2016/08/http.html) 对于协议有基本的认识和理解。

Go语言中提供了net/http包，通过http包可以很方便的搭建一个可以运行的web服务。

```go
package main

import(
	"fmt"
	"net/http"
	_ "strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)

	fmt.Fprintf(w, "hello")

}

func main() {
	http.HandleFunc("/", sayHelloName)

	err := http.ListenAndServe(":9190", nil)

	if err != nil {

		log.Fatal("ListenAndServe", err)
	}


}

```

通过http包```ListenAndServe```在端口9190监听。处理函数路径为"/"。





## 15.3 工作原理

对于http包中```ListenAndServe```整个过程，涉及到三个方面：

- 如何监听端口；
- 如何接收客户端请求？
- 如果分配handler。

底层处理是：初始化一个server对象，然后调用net.Listen("tcp", addr)，建立一个服务。Go语言中http包源代码为：

```go
func (srv *Server) Serve(l net.Listener) error {
		var tempDelay time.Duration // how long to sleep on accept failure

	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
	for {
		rw, err := l.Accept()
		if err != nil {
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
		connCtx := ctx
		if cc := srv.ConnContext; cc != nil {
			connCtx = cc(connCtx, rw)
			if connCtx == nil {
				panic("ConnContext returned nil")
			}
		}
		tempDelay = 0
		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew, runHooks) // before Serve can return
		go c.serve(connCtx)
	}
```

通过l.Accept()接收请求，然后创建一个srv.newConn(rw)新的连接，最后通过一个新的goroutine 服务该连接。完成高并发处理。



