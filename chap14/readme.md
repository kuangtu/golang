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

- 服务端：绑定端口进行监听，收到消息之后