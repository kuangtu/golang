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





## 参考文献

[墙裂推荐](https://www.devdungeon.com/content/working-files-go#write_bytes)