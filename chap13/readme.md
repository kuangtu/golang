#  第13章 标准库

本章对于golang标准库中常用包的使用进行说明。

## 13.1 flag包
flag包实现了命令行参数的解析。通过：```flag.String(), Bool(), Int()``` 定义不同类型的命令行参数。

示例：

```go
import "flag"
var nFlag = flag.Int("n", 1234, "help message for flag n")
```

定义了int类型的参数，**需要注意的是返回了int类型指针**，方法定义为：

```go
func Int(name string, value int, usage string) *int
```

参数：

- name，flag名称；
- value，默认值；
- usage，使用提示。

或者通过Var方法将参数绑定到某个变量：

```go
var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
```

定义了flag之后，调用```flag.Parse()``` 解析命令行参数到定义的flag中。

[示例程序](flag_example.go)：

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    wordPtr := flag.String("word", "foo", "a string")
    numberPtr := flag.Int("numb", 10, "an int")
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
    
    flag.Parse()
    
    fmt.Println("word", *wordPtr)
    fmt.Println("numb", *numberPtr)
    fmt.Println("svar", svar)
    
}
```

 

