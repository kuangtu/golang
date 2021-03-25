# 一、常用包的使用



## 1、命令行处理flag包

Golang 的标准库提供了 flag 包来处理命令行参数。



## 2、原子操作相关

atomic包通过提供低级别的原子内存原语来实现同步算法。

## 3、 strconv

标准的strconv实现了字符串和其他类型之间的相互转换。

### 3.1 数值转换

常用的是将字符串转为int，或者int转为字符串。

```go
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
```

## 4、容器

### 4.1 链表

Package list implements 实现了一个双向链表。

```go
package main

import (
    "fmt"
    "container/list"
)

func main() {
    fmt.Println("create list")
    //beego.Run()
    l := list.New()
    e4 := l.PushBack(4)
    eH := l.PushBack("H")
    
    //遍历
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
}
```

元素可以是任意的类型。



# 二、缓存相关



## 2.1、go-cach使用

go-cache是单机版本的内存key-value缓存库。任何对象都可以被存储，给定一个有效存续时间或者永久，cache可以被多个协程安全使用。

https://github.com/patrickmn/go-cache



## 2.1.1、基本使用

参照 [示例程序](go-cache/cache-test.go):

```go
package main


import (
    "fmt"
    "github.com/patrickmn/go-cache"
    "time"
)

func main() {
    //创建一个缓存
    c := cache.New(5*time.Minute, 10*time.Minute)
    //设置
    c.Set("foo", "bar", cache.DefaultExpiration)
    
    foo, found := c.Get("foo")
    if found {
        fmt.Println(foo)
    }

}
```



## 2.2 group-cached使用

[group-cached](https://github.com/golang/groupcache) 是一个分布高速缓存库，作为memcached的一个替代。【注：groupcached作者也是memcached的作者】。

### 2.2.1 基本介绍

和memcached类似，通过切片的方式由节点负责key。但是存在一些差异：

- group-cached不作为单独的服务器节点，大大减少了部署和配置的痛点，技术客户端也是服务端，它连接自身节点形成分布式缓存。
- 带有缓存填充机制。memcached 可能存在缓存雪崩的情况，因为数据未加载到缓存，或者缓存同一时间大面积的失效，从而导致所有请求都去查数据库，导致数据库CPU和内存负载过高，甚至宕机，group-cached协调缓存填充，**只有一个进程处理复制过程，然后复用该值到所有的调用者。**
- 不支持多个版本值。没有缓存过期时间、也没有明确的缓存收回。

但是发现没有example或者demo。

### 2.2.2 使用示例

参考网络[样例程序](https://sconedocs.github.io/groupcacheUseCase/)，基本使用方式为：

```go
package main

import (
    "errors"
    "flag"
    "log"
    "net/http"
    "strings"
    "context"
    "github.com/golang/groupcache"
)
//通过映射字面量初始化
var Store = map[string][]byte{
    "red":   []byte("#FF0000"),
    "green": []byte("#00FF00"),
    "blue":  []byte("#0000FF"),
}

var Group = groupcache.NewGroup("foobar", 64<<20, groupcache.GetterFunc(
    func(ctx context.Context, key string, dest groupcache.Sink) error {
        log.Println("looking up", key)
        v, ok := Store[key]
        if !ok {
            return errors.New("color not found")
        }
        dest.SetBytes(v)
        return nil
    },
))

func main() {
    addr := flag.String("addr", ":8080", "server address")
    peers := flag.String("pool", "http://localhost:8080", "server pool list")
    flag.Parse()
    http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
        color := r.FormValue("name")
        var b []byte
        err := Group.Get(nil, color, groupcache.AllocatingByteSliceSink(&b))
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.Write(b)
        w.Write([]byte{'\n'})
    })
    p := strings.Split(*peers, ",")
    pool := groupcache.NewHTTPPool(p[0])
    pool.Set(p...)
    http.ListenAndServe(*addr, nil)
}
```



