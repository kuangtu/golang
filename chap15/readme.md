# 第15章 web编程

本章对使用golang进行web编程进行整理。



## 15.1 概述

HTTP协议是互联网的基础协议，也是网页开发必备的基础知识，[可以参照阮一峰老师的文章](http://www.ruanyifeng.com/blog/2016/08/http.html) 对于协议有基本的认识和理解。

Go语言中提供了net/http包，通过http包可以很方便的搭建一个可以运行的web服务。

```go
package main

import(
	"fmt"
	"net/http"
	_ "strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)

	fmt.Fprintf(w, "hello")

}

func main() {
	http.HandleFunc("/", sayHelloName)

	err := http.ListenAndServe(":9190", nil)

	if err != nil {

		log.Fatal("ListenAndServe", err)
	}


}

```

通过http包```ListenAndServe```在端口9190监听。处理函数路径为"/"。



