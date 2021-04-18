# 第14章 Go网络编程

本章对Go语言网络编程内容进行整理，读者需要对网络基础知识有一定了解。



## 14.1 网络地址

对于TCP、UDP编程主要内容进行整理。



### 14.1.1 IP地址类型

Go语言中“net”包定义了多种类型，IP地址通过“byte slice”定义：

```go
type IP []byte
```

可以通过``ParseIP``方法解析[IP地址](ipaddr.go)：

```go
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
```

### 14.1.2 IPAddr类型

部分函数返回了IPAddr指针，是一个简单的结构体包含了一个IP。

```go
type IPAddr {
    ip IP
}
```

常用于DNS查询获得IP地址：

```go
func	ResolveIPAddr(net,	addr	string)	(*IPAddr,	os.Error)
```

[解析IP地址](resolveip.go)：

```go
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
    
```



### 14.1.3 主机查询

``ResolveIPAddr`` DNS查询返回一个IP地址。如果主机有多个网络接口配置多个IP地址，有不同的主机名称，通过如下函数查询：

```go
func	LookupHost(name	string)	(addrs	[]string,	err	os.Error)
```

规范主机查询：

```go
func	LookupCNAME(name	string)	(cname	string,	err	os.Error)
```

[查询多个IP地址](lookuphost.go):

```go
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
    
```

**``LookUpHost`` 返回的是多个字符串，而不是``IPAddress``类型。**



## 14.2 网络服务

主机运行网络服务，通常等待请求并进行回复。



### 14.2.1 端口

网络服务通常在特定端口提供服务，*nix环境中常用端口存放在：``/etc/services``。通过``LookupPort``函数查询：

```go
func	LookupPort(network,	service	string)	(port	int,	err	os.Error)
```



``network`` 参数为类型类型：tcp、udp。``service``为服务类型。[端口查询类型](lookupport.go):

```go
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
    
```



### 14.2.2 TCPAddr地址

类型``TCPAddr``包含了IP地址和端口的结构体：

```go
type	TCPAddr	struct	{ 
    IP		IP
	Port	int 
}
```

通过方法``ResolveTCPAddr``创建TCPAddr：

```go
func	ResolveTCPAddr(net,	addr	string)	(*TCPAddr,	os.Error)
```

``net``是tcp、udp等，addr是主机名称或者IP以及通过“:"连接的端口号，比如：www.google.com:80、127.0.0.1:223。



## 14.3 TCP Sockets

服务端、客户端之间通信：

- 服务端：绑定端口进行监听，收到消息处理之后回复消息。

主要通过``net.TCPConn``类型完成客户端-服务端之间的双工通信，主要有两个方法：

```go
func	(c	*TCPConn)	Write(b	[]byte)	(n	int,	err	os.Error)
func	(c	*TCPConn)	Read(b	[]byte)	(n	int,	err	os.Error)
```



### 14.3.1 TCP客户端

客户端发起连接，成功后返回TCPConn进行通信。TCP连接通过``DialTCP``函数：

```go
func	DialTCP(net	string,	laddr,	raddr	*TCPAddr)	(c	*TCPConn,	err	os.Error)
```

其中的``laddr``是本地地址，通常设置为nil（**如果是本机多网络地址？**）,raddr是提供服务的远程地址，``net``字符串是网络类型：“tcp"、"tcp4"、"tcp6"。

[访问web服务](getheadinfo.go):

```go
package main

import(
    "net"
    "os"
    "fmt"
    "io/ioutil"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "usage: %s host:port\n", os.Args[0])
        os.Exit(1)
    }
    
    service := os.Args[1]
    
    //解析TCP地址及端口
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    
    //发起TCP连接
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)
    
    //向服务端发送head命令
    _, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)
    
    result, err := conn.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))
    
    os.Exit(0)
    
}

func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }

```

需要注意一些异常的处理：

- 地址语法错误
- 连接远程服务失败
- 即时连接成功，读写的时候也可能以为网络中断、超时等出现异常。

TODO，读写比较。



### 14.3.2 DayTime服务示例

daytime服务是在RFC 867中定义的网络服务，端口为：13，可以通过TCP或者UDP访问。

基本流程：

- 服务端注册端口后监听，然后阻塞在``accept``操作，等待客户端接入。
- 当客户端连接之后，accept调用返回连接对象。

主要函数调用如下：

```go
func	ListenTCP(net	string,	laddr	*TCPAddr)	(l	*TCPListener,	err	os.Error) func	(l	*TCPListener)	Accept()	(c	Conn,	err	os.Error)
```



[示例程序](daytimeserver.go):

```go
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
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    
    //监听端口
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    
    for {
        conn, err := listener.Accept()
        
        if err != nil {
            conntinue
        }
        
        daytime := time.Now().String()
        conn.Write([]byte(daytime))
        conn.Close()
    }
    
}

func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }

```

  

这个服务每次只能接收一个客户端。可以通过“多线程”的方式服务服务多个客户端，[程序示例](thdechoserver.go):

```go
package main

import(
    "net"
    "os"
    "fmt"
)

func main() {

    service := ":1200"

    //解析TCP地址及端口
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    
    //监听端口
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    
    for {
        conn, err := listener.Accept()
        
        if err != nil {
            conntinue
        }
        
        go handleClient(conn)
    }
    
}

func handleClient(conn net.Conn) {
    defer conn.Close()
    
    var buf [512]byte
    for {
        n, err := conn.Read(buf[0:])
        if err != nil {
            return
        }
        
        //write data，收到了
        _, err2 := conn.Write(buf[0:n])
        if err2 != nil {
            return 
        }
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }


```

通过关键字``go``采用go-coroutine方式执行，完成“多线程”服务。