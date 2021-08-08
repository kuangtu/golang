# 第16章 文件IO

程序处理过程中，通常会使用文件读写，用于存储数据。本章对于文件基本操作和文件读写进行说明。

## 16.1 文件基本操作

### 16.1.1 文件创建

通过```create```函数创建文件，如果文件已经存在，则截断该文件。如果不存在，按照权限0666进行创建。‘0666’权限表示其他用户也可以访问。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("create file demo")
	name := "test.txt"
	fp, err := os.Create(name)
	if err != nil {
		fmt.Println("create file error:", err)
	}
	fp.Close()
}

```

### 16.1.2 重命名文件

通过```rename```函数重命名文件：

```go
	src := "srctest.txt"
	dst := "dsttest.txt"
	renerr := os.Rename(src, dst)
	if renerr != nil {
		fmt.Println("rename error:", err)
	}
```

如果目标文件存在且不是目录，会替换该文件。



## 16.2 Buffering IO
Golang语言提供了bufio包，以缓存方式累积结果后进行IO操作。通过减少了系统调用增加程序速度。

### 16.2.1 bufio.Writer
可以通过bufio.Writer方法在IO写入之前累积数据到buffer中。
- buffer满了
如果buffer满了，立即进行写入操作。
- buffer还有空间
不会尝试写入，直到通过Flush()方法告知。
- 超过buffer容量
跳过buffer直接写入。

## 参考文献

[墙裂推荐](https://www.devdungeon.com/content/working-files-go#write_bytes)
[bufio](https://www.educative.io/edpresso/how-to-read-and-write-with-golang-bufio)
[文件操作](https://golangbot.com/write-files/)