# 第11章 测试

我们说测试的时候一般是指自动化测试，也就是写一些小的程序用来检测被测试代码（产品代码）的行为和预期的一样，这些通常都是精心设计的执行某些特定的功能或者是通过随机性的输入待验证边界的处理。

Go语言的测试技术是相对低级的。它依赖一个go test测试命令和一组按照约定方式编写的测试函数，测试命令可以运行这些测试函数。编写相对轻量级的纯测试代码是有效的，而且它很容易延伸到基准测试和示例文档。



## 11.1 go test

go test命令是一个按照一定的约定和组织来测试代码的程序。在包目录内，所有以`_test.go`为后缀名的源文件在执行go build时不会被构建成包的一部分，它们是go test测试的一部分。

在`*_test.go`文件中，有三种类型的函数：测试函数、基准测试（benchmark）函数、示例函数。

- 一个测试函数是以Test为函数名前缀的函数，用于测试程序的一些逻辑行为是否正确；go test命令会调用这些测试函数并报告测试结果是PASS或FAIL。
- 基准测试函数是以Benchmark为函数名前缀的函数，它们用于衡量一些函数的性能；go test命令会多次运行基准测试函数以计算一个平均的执行时间。
- 示例函数是以Example为函数名前缀的函数，提供一个由编译器保证正确性的示例文档。



## 11.2 测试函数

每个测试函数必须导入testing包。测试函数有如下的签名：

```go
func TestName(t *testing.T) {
    // ...
}
```

测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头：

```go
func TestSin(t *testing.T) { /* ... */ }
func TestCos(t *testing.T) { /* ... */ }
func TestLog(t *testing.T) { /* ... */ }
```

回文检查程序：

```go
package word

func IsPalindrome(s string) bool {
    for i := range s {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}
```

这个实现对于一个字符串是否是回文字符串前后重复测试了两次；

测试程序：

```go
package word

import "testing"

func TestPalindrome(t *testing.T) {
    if !IsPalindrome("detartrated") {
        t.Error(`IsPalindrome("detartrated") = false`)
    }
    if !IsPalindrome("kayak") {
        t.Error(`IsPalindrome("kayak") = false`)
    }
}

func TestNonPalindrome(t *testing.T) {
    if IsPalindrome("palindrome") {
        t.Error(`IsPalindrome("palindrome") = true`)
    }
}
```

`go test`命令如果没有参数指定包那么将默认采用当前目录对应的包。



## 11.3 测试覆盖率

就其性质而言，测试不可能是完整的。计算机科学家Edsger Dijkstra曾说过：“测试能证明缺陷存在，而无法证明没有缺陷。”再多的测试也不能证明一个程序没有BUG。在最好的情况下，测试可以增强我们的信心：代码在很多重要场景下是可以正常工作的。

`go tool`命令运行Go工具链的底层可执行程序。这些底层可执行程序放在GOROOT/pkg/tool/GOROOT/pkg/tool/{GOOS}_${GOARCH}目录。



## 11.4 基准测试

基准测试是测量一个程序在固定工作负载下的性能。在Go语言中，基准测试函数和普通测试函数写法类似，但是以Benchmark为前缀名，并且带有一个`*testing.B`类型的参数；`*testing.B`参数除了提供和`*testing.T`类似的方法，还有额外一些和性能测量相关的方法。它还提供了一个整数N，用于指定操作执行的循环次数。

## 单元测试
单元测试是用来测试程序的一部分代码或者一组代码的函数。
- 正向路径测试，在正常执行的情况下，保障代码不产生错误的测试；
- 负向路径测试，保证代码不仅会产生错误，而且是预期的错误。
Go语言中有几种写单元测试的方法：
- 基础测试，只使用一组参数和结果来测试一段代码；
- 表组测试，也会测试一段代码，但是会使用多组参数和结果进行测试；
- 用一些方法mock测试代码需要使用到的外部资源，比如：数据库或者网络服务器。


``` go
// 这个示例程序展示如何写基础单元测试
package listing01

import ( 
   "net/http"
   "testing"
) 

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload确认  http包的  Get函数可以下载内容
func TestDownload(t *testing.T) { 
   url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
   statusCode := 200

   t.Log("Given the need to test downloading content.")
   { 
       t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
           url, statusCode)
       { 
           resp, err := http.Get(url)
           if err != nil { 
               t.Fatal("\t\tShould be able to make the Get call.",
                   ballotX, err)
           } 
           t.Log("\t\tShould be able to make the Get call.",
               checkMark)

           defer resp.Body.Close()

           if resp.StatusCode == statusCode { 
               t.Logf("\t\tShould receive a \"%d\" status. %v",
                   statusCode, checkMark)
           } else { 
               t.Errorf("\t\tShould receive a \"%d\" status. %v %v",
                   statusCode, ballotX, resp.StatusCode)
           } 
       } 
   } 
} 
```



