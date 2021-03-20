# 一、常用包的使用



## 1、命令行处理flag包

Golang 的标准库提供了 flag 包来处理命令行参数。



## 2、原子操作相关





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

