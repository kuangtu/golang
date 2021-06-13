# 对于golang项目按照分类进行整理。



## 1、数据库相关

[Badger: 快速Key-Value数据库](https://github.com/dgraph-io/badger)

基本介绍：



相关资料：

[Introducing Badger: A fast key-value store written purely in Go](https://dgraph.io/blog/post/badger/) Dgraph网站的介绍，对于Badger的开发背景等进行说明。

[Badger: Fast Key-Value DB in Go](http://bos.itdks.com/a121f6647d6042989fb9e76fa40a03f4.pdf) Gopher China 2018演讲。



# 2、网络相关

[tcp网络代理](https://github.com/jpillora/go-tcp-proxy)

基本介绍：

```go
$ tcp-proxy --help
Usage of tcp-proxy:
  -c: output ansi colors
  -h: output hex
  -l="localhost:9999": local address
  -n: disable nagles algorithm
  -r="localhost:80": remote address
  -match="": match regex (in the form 'regex')
  -replace="": replace regex (in the form 'regex~replacer')
  -v: display server actions
  -vv: display server actions and all tcp data
```

-l参数，本地地址和端口；

-r参数， 远程地址。

HTTP协议基于TCP，可以将go tcp-proxy作为HTTP代理：

```go
$ tcp-proxy -r echo.jpillora.com:80
Proxying from localhost:9999 to echo.jpillora.com:80
```







