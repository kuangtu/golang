# 第十章 包和工具

Go语言有超过100个的标准包，标准库为大多数的程序提供了必要的基础构件。

```go
go list std 
```

## 10.1 包简介

任何包系统设计的目的都是为了简化大型程序的设计和维护工作，通过将一组相关的特性放进一个独立的单元以便于理解和更新，在每个单元更新的同时保持和程序中其它单元的相对独立性。这种模块化的特性允许每个包可以被其它的不同项目共享和重用，在项目范围内、甚至全球范围统一的分发和复用。

每个包一般都定义了一个不同的名字空间用于它内部的每个标识符的访问。每个名字空间关联到一个特定的包，让我们给类型、函数等选择简短明了的名字，这样可以在使用它们的时候减少和其它部分名字的冲突。



## 10.2 导入路径

每个包是由一个全局唯一的字符串所标识的导入路径定位。出现在import语句中的导入路径也是字符串。

```go
import (
    "fmt"
    "math/rand"
    "encoding/json"

    "golang.org/x/net/html"

    "github.com/go-sql-driver/mysql"
)
```

如果你计划分享或发布包，那么导入路径最好是全球唯一的。为了避免冲突，所有非标准库包的导入路径建议以所在组织的互联网域名为前缀；而且这样也有利于包的检索。例如，上面的import语句导入了Go团队维护的HTML解析器和一个流行的第三方维护的MySQL驱动。

## 10.3 包声明

在每个Go语言源文件的开头都必须有包声明语句。包声明语句的主要目的是确定当前包被其它包导入时默认的标识符（也称为包名）。

例如，math/rand包的每个源文件的开头都包含`package rand`包声明语句，所以当你导入这个包，你就可以用rand.Int、rand.Float64类似的方式访问包的成员。

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println(rand.Int())
}
```



## 10.4 导入声明

可以在一个Go语言源文件包声明语句之后，其它非导入声明语句之前，包含零到多个导入包声明语句。每个导入声明可以单独指定一个导入路径，也可以通过圆括号同时导入多个导入路径。下面两个导入形式是等价的，但是第二种形式更为常见。

```go
import "fmt"
import "os"

import (
    "fmt"
    "os"
)
```

果我们想同时导入两个有着名字相同的包，例如math/rand包和crypto/rand包，那么导入声明必须至少为一个同名包指定一个新的包名以避免冲突。这叫做**导入包的重命名**。

```go
import (
    "crypto/rand"
    mrand "math/rand" // alternative name mrand avoids conflict
)
```



## 10. 5 包的匿名导入

如果只是导入一个包而并不使用导入的包将会导致一个编译错误。

但是有时候我们只是想利用导入包而产生的副作用：它会计算包级变量的初始化表达式和执行导入包的init初始化函数。这时候我们需要抑制“unused import”编译错误，我们可以用下划线`_`来重命名导入的包。像往常一样，下划线`_`为空白标识符，并不能被访问。

```go
import _ "image/png" // register PNG decoder
```



## 10.6  包和命名

当创建一个包，一般要用短小的包名，但也不能太短导致难以理解。标准库中最常用的包有bufio、bytes、flag、fmt、http、io、json、os、sort、sync和time等包。

尽可能让命名有描述性且无歧义。

包名一般采用单数的形式。



## 10.7 工具

本章剩下的部分将讨论Go语言工具箱的具体功能，包括如何下载、格式化、构建、测试和安装Go语言编写的程序。

```go
$ go
...
    build            compile packages and dependencies
    clean            remove object files
    doc              show documentation for package or symbol
    env              print Go environment information
    fmt              run gofmt on package sources
    get              download and install packages and dependencies
    install          compile and install packages and dependencies
    list             list packages
    run              compile and run Go program
    test             test packages
    version          print Go version
    vet              run go tool vet on packages

Use "go help [command]" for more information about a command.
...
```

## 10.8 包的下载安装

go程序编写过程中会用到第三方库，很多使用了github。在go build过程中会下载第三方库。默认的go proxy代理很慢，可以进行修改。参考[链接](https://goproxy.cn/)。

```bash
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
```

可以避免网络超时的情况。

通过 go env，查看相应的配置：

![go_env环境](C:\moxuansheng\workspace\golang\jpg\go_env环境.png)

可以看到相应的包缓存路径。有些github包安装如下。

![包的位置](C:\moxuansheng\workspace\golang\jpg\包的位置.png)



https://zhuanlan.zhihu.com/p/60703832

拜拜了，GOPATH君！新版本Golang的包管理入门教程

## 10.9包管理

（1）GOPATH使用

是GO开发环境所设置的一个变量。历史版本的 go 语言开发时，需要将代码放在 GOPATH 目录的 src 文件夹下。go get 命令获取依赖，也会自动下载到 GOPATH 的 src 下。GOPATH设置如下：

![gopath设置](C:\moxuansheng\workspace\golang\jpg\gopath设置.png)

通过go get进行包下载：

![goget](C:\moxuansheng\workspace\golang\jpg\goget.png)

可以看到下载的数据包在GOPATH路径中：

![goget_envpath](C:\moxuansheng\workspace\golang\jpg\goget_envpath.png)

（2）存在的问题

GOPATH 模式下，go get 命令使用时，没有版本选择机制，拉下来的依赖代码都会默认当前最新版本，而且如果当项目 A 和项目 B 分别依赖项目 C 的两个不兼容版本时， GOPATH 路径下只有一个版本的 C 将无法同时满足 A 和 B 的依赖需求。这可以说是一个很大的缺陷了，因而 Go1.13 起，官方就不再推荐使用 GOPATH 的模式了。

(3) Go modules 

基本思路是**为每个项目单独维护一份对应版本依赖的拷贝**。主要涉及到如下几个配置：

```go
GO111MODULE="auto"
GOPROXY="https://goproxy.io,direct"
GONOPROXY=""
GOSUMDB="sum.golang.org"
GONOSUMDB=""
GOPRIVATE=""
```

是否开启go module的开关，可用参数值如下：

|       值       |                             含义                             |
| :------------: | :----------------------------------------------------------: |
|      auto      | 在 GOPATH/src 之外，将自动使用 Go Modules 模式。否则还是用 GOPATH 模式。目前在最新的 Go1.14 中是默认值。 |
|       on       | 启用 Go modules，将不使用 GOPATH，推荐设置，将会是未来版本中的默认值。 |
| off 或者不设置 | Go 将使用 GOPATH 和沿用老的 vendor 机制，禁用 Go modules。不推荐设置。 |



【参考文献】https://www.infoq.cn/article/xyjhjja87y7pvu1iwhz3

（4）go mod init命令

项目初始化：

```go
go mod init github/kuangtu/groupcache_test
```

可以看到创建了文件go.mod：

![goinit](C:\moxuansheng\workspace\golang\jpg\goinit.png)

文件内容如下：

```
module github/kuangtu/groupcache_test

go 1.16
```

它记录了当前项目的模块信息，每一行都以一个关键词开头。在该目录下执行go get，下载依赖包：

![go get pkg](C:\moxuansheng\workspace\golang\jpg\go get pkg.png)

再次查看go.mod文件：

```
module github/kuangtu/groupcache_test

go 1.16

require (
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3 // indirect
)
```

go.sum文件：

下载依赖包有可能被恶意篡改，以及缓存在本地的依赖包也有被篡改的可能，单单一个 go.mod 文件并不能保证一致性构建，Go 开发团队在引入 go.mod 的同时也引入了 go.sum 文件，用于记录每个依赖包的哈希值（SHA-256 算法），在 build 时，如果本地的依赖包 hash 值与 go.sum 文件中记录得不一致，则会拒绝 build。